package service

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

var FCMService FCM

type FCM interface {
	SendMessage(context.Context, *messaging.Message) error
}

func New(fcm *firebase.App) {
	FCMService = &fcmService{
		fcm: fcm,
	}
}

type fcmService struct {
	fcm *firebase.App
}

// SendMessage implements FCM.
func (f *fcmService) SendMessage(ctx context.Context, message *messaging.Message) error {
	msg, err := f.fcm.Messaging(ctx)
	if err != nil {
		return err
	}

	_, err = msg.Send(ctx, message)

	return err
}
