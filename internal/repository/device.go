package repository

import (
	"context"

	"notifsys/abstraction"
	"notifsys/internal/dto"
	"notifsys/internal/model"
	"notifsys/internal/repository/interfaces"

	"github.com/uptrace/bun"
)

func NewDevice(db *bun.DB) interfaces.Device {
	Device := &deviceRepository{
		DB: db,
	}

	return Device
}

type deviceRepository struct {
	*bun.DB
}

// Create implements DeviceRepository.
func (r *deviceRepository) Create(ctx context.Context, payload *model.Device) (*model.Device, error) {
	_, err := r.DB.NewInsert().Model(payload).Exec(ctx)

	return payload, err
}

// Find implements DeviceRepository.
func (r *deviceRepository) Find(ctx context.Context, f *dto.DeviceFilter, p *abstraction.Pagination) ([]*model.Device, error) {
	panic("unimplemented")
}

// FindOne implements DeviceRepository.
func (r *deviceRepository) FindOne(ctx context.Context, f *dto.DeviceFilter) (*model.Device, error) {
	result := &model.Device{}

	err := r.NewSelect().Model(result).ApplyQueryBuilder(f.Apply).Scan(ctx, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
