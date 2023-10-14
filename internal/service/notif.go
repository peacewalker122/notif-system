package service

import (
	"context"
	"log"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/repository/interfaces"

	"firebase.google.com/go/messaging"
	"github.com/uptrace/bun"
)

func NewNotif(f *factory.Factory) *Notif {
	return &Notif{
		DB:   f.DB,
		User: f.User,
	}
}

type Notif struct {
	DB   *bun.DB
	User interfaces.User
}

// Create implements Service.
func (s *Notif) Create(ctx context.Context, payload *dto.NotifRequest) error {
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

	log.Printf("data: %+v\n", devicestoken)

	err = FCMService.SendMessage(ctx, &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: payload.Message["title"],
			Body:  payload.Message["body"],
		},
		Tokens: devicestoken,
	})

	return err
}
