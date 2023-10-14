package interfaces

import (
	"context"

	"notifsys/internal/dto"
)

type Notif interface {
	Create(ctx context.Context, payload *dto.NotifRequest) error
}
