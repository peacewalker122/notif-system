package repository

import (
	"context"

	"notifsys/abstraction"
	"notifsys/internal/dto"
	"notifsys/internal/model"

	"github.com/uptrace/bun"
)

var (
	_    UserRepository = (*userRepository)(nil)

	User UserRepository
)

type UserRepository interface {
	Create(ctx context.Context, payload *model.User) (*model.User, error)
	Find(ctx context.Context, f *dto.UserFilter, p *abstraction.Pagination) ([]*model.User, error)
	FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error)
}

func NewUser(db *bun.DB) {
	User = &userRepository{
		DB: db,
	}
}

type userRepository struct {
	*bun.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(ctx context.Context, payload *model.User) (*model.User, error) {
	_, err := u.DB.NewInsert().Model(payload).Exec(ctx)

	return payload, err
}

// Find implements UserRepository.
func (u *userRepository) Find(ctx context.Context, f *dto.UserFilter, p *abstraction.Pagination) ([]*model.User, error) {
	result := make([]*model.User, 0)

	err := u.NewSelect().Model((*model.User)(nil)).ApplyQueryBuilder(f.Apply).Scan(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error) {
	result := &model.User{}

	err := u.NewSelect().Model(result).ApplyQueryBuilder(f.Apply).Scan(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
