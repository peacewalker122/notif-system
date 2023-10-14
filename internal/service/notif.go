package service

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/repository/interfaces"

	"firebase.google.com/go/messaging"
	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewNotif(f *factory.Factory) *Notif {
	return &Notif{
		DB:   f.DB,
		User: f.User,
		FCM:  f.FCM,
	}
}

type Notif struct {
	DB   *bun.DB
	User interfaces.User
	FCM  interfaces.FCM
}

// Create implements Service.
func (s *Notif) Create(ctx context.Context, payload *dto.NotifRequest) error {
	// Get the active span from the context.
	span := trace.SpanFromContext(ctx)

	if span.IsRecording() {
		span.AddEvent("Create Notif", trace.WithAttributes(
			attribute.IntSlice("user_id", payload.UserID),
		))
	}

	data, err := s.User.Find(ctx, &dto.UserFilter{
		ID:              payload.UserID,
		WithDeviceToken: true,
	}, nil)
	if err != nil {
		return err
	}

	devicestoken := make([]string, 0)

	for _, v := range data {
		if v.DeviceToken != nil {
			devicestoken = append(devicestoken, *v.DeviceToken)
		}
	}

	err = s.FCM.SendMessage(ctx, &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: payload.Message["title"],
			Body:  payload.Message["body"],
		},
		Tokens: devicestoken,
	})

	return err
}
