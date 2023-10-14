package repository

import (
	"context"

	"notifsys/abstraction"
	"notifsys/internal/dto"
	"notifsys/internal/model"
	"notifsys/internal/repository/interfaces"

	"github.com/uptrace/bun"
)

func NewUser(db *bun.DB) interfaces.User {
	User := &userRepository{
		DB: db,
	}

	return User
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

	err := u.NewSelect().Model((*model.User)(nil)).
		Apply(func(sq *bun.SelectQuery) *bun.SelectQuery {
			if f != nil {
				f.Apply(sq)
			}

			return sq
		}).Scan(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindOne(ctx context.Context, f *dto.UserFilter) (*model.User, error) {
	result := &model.User{}

	err := u.NewSelect().
		Model((*model.User)(nil)).
		Apply(func(sq *bun.SelectQuery) *bun.SelectQuery {
			if f != nil {
				f.Apply(sq)
			}

			return sq
		}).Scan(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
