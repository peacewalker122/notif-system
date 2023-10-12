package service

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/repository"

	"firebase.google.com/go/messaging"
	"github.com/uptrace/bun"
)

type Notif interface {
	Create(ctx context.Context, payload *dto.NotifRequest) error
}

func NewNotif(DB *bun.DB) Notif {
	return &notif{
		DB: DB,
	}
}

type notif struct {
	DB *bun.DB
}

// Create implements Service.
func (s *notif) Create(ctx context.Context, payload *dto.NotifRequest) error {
	data, err := repository.User.Find(ctx, &dto.UserFilter{
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

	err = FCMService.SendMessage(ctx, &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: payload.Message["title"],
			Body:  payload.Message["body"],
		},
		Tokens: devicestoken,
	})

	return err
}
