package service

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/factory"
	"notifsys/internal/model"
	"notifsys/internal/repository/interfaces"
	"notifsys/pkg/bcrypt"
	"notifsys/pkg/trx"

	"github.com/uptrace/bun"
)

func NewUser(f *factory.Factory) *User {
	return &User{
		DB:   f.DB,
		User: f.User,
	}
}

type User struct {
	DB     *bun.DB
	User   interfaces.User
	Device interfaces.Device
}

// Create implements Service.
func (s *User) Create(ctx context.Context, payload *dto.SignupRequest) (*model.User, error) {
	var data *model.User

	pass, err := bcrypt.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	payload.Password = pass

	err = trx.New(s.DB).Run(ctx, func(ctx context.Context) error {
		var err error

		data, err = s.User.Create(ctx, &model.User{
			Username: payload.Name,
			Email:    payload.Email,
			Password: payload.Password,
		})
		if err != nil {
			return err
		}

		_, err = s.Device.Create(ctx, &model.Device{
			UserID:      &data.ID,
			DeviceToken: payload.DeviceToken,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindOne implements Service.
func (s *User) FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error) {
	return s.User.FindOne(ctx, f)
}
