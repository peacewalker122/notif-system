package interfaces

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/model"
)

type User interface {
	Create(ctx context.Context, payload *dto.SignupRequest) (*model.User, error)
	FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error)
}
