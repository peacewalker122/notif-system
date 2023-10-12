package factory

import (
	"notifsys/internal/repository"

	"github.com/uptrace/bun"
)

func Run(db *bun.DB) {
	repository.NewUser(db)
	repository.NewDevice(db)
}
