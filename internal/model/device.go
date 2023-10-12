package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type Device struct {
	bun.BaseModel `bun:"table:devices,alias:d"`

	ID          int        `bun:"id,pk"`
	UserID      *int       `bun:"user_id"`
	DeviceToken *string    `bun:"device_token"`
	LastLoginAt *time.Time `bun:"last_login_at"`
}

// BeforeAppendModel implements schema.BeforeAppendModelHook.
func (d *Device) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	tnow := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		d.LastLoginAt = &tnow
	case *bun.UpdateQuery:
		d.LastLoginAt = &tnow
	}

	return nil
}

func (*Device) TableName() string {
	return "devices"
}

var _ bun.BeforeAppendModelHook = (*Device)(nil)
