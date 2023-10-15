package cron

import (
	"context"
	"encoding/json"
	"log/slog"
	"runtime"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/repository/interfaces"
	"notifsys/pkg/amqp"
	"notifsys/pkg/tracer"

	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/trace"

	"firebase.google.com/go/messaging"
	"github.com/uptrace/bun"
	"github.com/uptrace/uptrace-go/uptrace"
)

type Service interface {
	Notification(context.Context)
}

func NewService(f *factory.Factory) *service {
	return &service{
		db:   f.DB,
		amqp: f.AMQP,
		fcm:  f.FCM,
		user: f.User,
		done: f.Done,
	}
}

type service struct {
	db   *bun.DB
	amqp *amqp.AMQP
	done chan struct{}

	fcm  interfaces.FCM
	user interfaces.User
}

func (s *service) Notification(ctx context.Context) {
	msgs, done, err := s.amqp.Consume(ctx, "notif")
	if err != nil {
		slog.Error("amqp-consume", err)
		return
	}
	defer done()

	go func() {
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return // Channel is closed
				}
				go func(msg amqp091.Delivery) {
					var err error
					ctx, trc := tracer.Trace.Start(context.Background(), "amqp-consume")
					defer func() {
						if r := recover(); r != nil {
							buf := make([]byte, 4096)
							runtime.Stack(buf, false)
							// Log the panic here or take any other appropriate action.
							slog.Error("Recover from panic:", string(buf))
						}

						if err != nil {
							slog.Error("[CRON] sendtotrace", err)
							trc.RecordError(err, trace.WithStackTrace(true))
						}
						trc.End()
						slog.Info("[TRACE] ", uptrace.TraceURL(trc))
					}()

					var mqtbody dto.NotifRequest
					if err := json.Unmarshal(msg.Body, &mqtbody); err != nil {
						slog.Error("[CRON] amqp-consume", err)
						return
					}

					data, err := s.user.Find(ctx, &dto.UserFilter{
						ID:              mqtbody.UserID,
						WithDeviceToken: true,
					}, nil)
					if err != nil {
						slog.Error("[CRON] amqp-consume", err)
						return
					}

					devicetokens := make([]string, len(data))
					for i, v := range data {
						devicetokens[i] = *v.DeviceToken
					}

					err = s.fcm.SendMessage(ctx, &messaging.MulticastMessage{
						Notification: &messaging.Notification{
							Title: mqtbody.Message["title"],
							Body:  mqtbody.Message["body"],
						},
						Tokens: devicetokens,
					})
					if err != nil {
						slog.Error("[CRON] amqp-consume", err)
						return
					}
				}(msg)
			case <-s.done:
				slog.Info("amqp-consume-done")
				return
			}
		}
	}()

	<-s.done
}
