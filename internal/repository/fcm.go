package repository

import (
	"context"

	"notifsys/internal/repository/interfaces"
	"notifsys/pkg/tracer"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func New(fcm *firebase.App) interfaces.FCM {
	fcmstruct := &fcmService{
		fcm: fcm,
	}

	return fcmstruct
}

type fcmService struct {
	fcm *firebase.App
}

// SendMessage implements FCM.
func (f *fcmService) SendMessage(ctx context.Context, message *messaging.MulticastMessage) error {
	ctx, trc := tracer.Trace.Start(ctx, "SendMessage")
	defer trc.End()

	msg, err := f.fcm.Messaging(ctx)
	if err != nil {
		return err
	}

	_, err = msg.SendMulticast(ctx, message)

	return err
}
