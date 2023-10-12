package dto

import (
	"github.com/uptrace/bun"
)

type UserFilter struct {
	ID    []int   `json:"id"`
	Email *string `json:"email"`

	WithDeviceToken bool `json:"-"`
}

func (u *UserFilter) Apply(db *bun.SelectQuery) {
	if len(u.ID) > 0 {
		db = db.Where("u.id IN (?)", bun.In(u.ID))
	}
	if u.Email != nil {
		db = db.Where("u.email = ?", *u.Email)
	}
	if u.WithDeviceToken {
		db = db.Join("JOIN devices ON devices.user_id = u.id")
		db.ColumnExpr("u.*,devices.device_token as device_token")
	}
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignupRequest struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Phone    string `json:"phone,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`

	DeviceToken *string `json:"device_token,omitempty"`
}
