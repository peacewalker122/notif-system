package service

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

var FCMService FCM

type FCM interface {
	SendMessage(context.Context, *messaging.MulticastMessage) error
}

func New(fcm *firebase.App) FCM {
	fcmstruct := &fcmService{
		fcm: fcm,
	}

	FCMService = fcmstruct

	return fcmstruct
}

type fcmService struct {
	fcm *firebase.App
}

// SendMessage implements FCM.
func (f *fcmService) SendMessage(ctx context.Context, message *messaging.MulticastMessage) error {
	msg, err := f.fcm.Messaging(ctx)
	if err != nil {
		return err
	}

	_, err = msg.SendMulticast(ctx, message)

	return err
}
