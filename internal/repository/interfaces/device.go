package interfaces

import (
	"context"

	"notifsys/internal/dto"
	"notifsys/internal/model"
)

type Device interface {
	Create(ctx context.Context, payload *model.Device) (*model.Device, error)
	FindOne(ctx context.Context, f *dto.DeviceFilter) (*model.Device, error)
}
