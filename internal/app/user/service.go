package user

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/model"
	"notifsys/internal/repository"
	"notifsys/pkg/bcrypt"
	"notifsys/pkg/trx"

	"github.com/uptrace/bun"
)

type Service interface {
	Create(ctx context.Context, payload *dto.SignupRequest) (*model.User, error)
	FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error)
}

func NewService(DB *bun.DB) *service {
	return &service{
		DB: DB,
	}
}

type service struct {
	DB *bun.DB
}

// Create implements Service.
func (s *service) Create(ctx context.Context, payload *dto.SignupRequest) (*model.User, error) {
	var data *model.User

	pass, err := bcrypt.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	payload.Password = pass

	err = trx.New(s.DB).Run(ctx, func(ctx context.Context) error {
		var err error

		data, err = repository.User.Create(ctx, &model.User{
			Username: payload.Name,
			Email:    payload.Email,
			Password: payload.Password,
		})
		if err != nil {
			return err
		}

		_, err = repository.Device.Create(ctx, &model.Device{
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
func (s *service) FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error) {
	return repository.User.FindOne(ctx, f)
}

var _ Service = (*service)(nil)
