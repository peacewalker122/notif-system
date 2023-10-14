package interfaces

import (
	"context"

	"notifsys/abstraction"
	"notifsys/internal/dto"
	"notifsys/internal/model"
)

type User interface {
	Create(ctx context.Context, payload *model.User) (*model.User, error)
	Find(ctx context.Context, f *dto.UserFilter, p *abstraction.Pagination) ([]*model.User, error)
	FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error)
}
