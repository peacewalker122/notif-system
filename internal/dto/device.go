package dto

import "github.com/uptrace/bun"

type DeviceFilter struct {
	ID     []int `json:"id"`
	UserID []int `json:"user_id"`
}

func (r *DeviceFilter) Apply(db bun.QueryBuilder) bun.QueryBuilder {
	if len(r.ID) > 0 {
		db = db.Where("id IN (?)", bun.In(r.ID))
	}
	if len(r.UserID) > 0 {
		db = db.Where("user_id IN (?)", bun.In(r.UserID))
	}

	return db
}
