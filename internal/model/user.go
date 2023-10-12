package model

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

// User represents a user in the database.
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int        `bun:"id,pk,autoincrement"`
	Email         string     `bun:"email"`
	Username      string     `bun:"username"`
	Password      string     `bun:"password"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     *time.Time `bun:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete"`

	DeviceToken *string `bun:"device_token,scanonly"`
}

// BeforeAppendModel implements schema.BeforeAppendModelHook.
func (u *User) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	tnow := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		u.CreatedAt = tnow
	case *bun.UpdateQuery:
		u.UpdatedAt = &tnow
	}

	return nil
}

func (*User) TableName() string {
	return "users"
}

var _ bun.BeforeAppendModelHook = (*User)(nil)
