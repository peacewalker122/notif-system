package factory

import (
	"notifsys/internal/repository"
	"notifsys/internal/service"
	"notifsys/pkg/fcm"

	"github.com/uptrace/bun"
)

func Run(db *bun.DB) {
	service.New(fcm.APP)

	repository.NewUser(db)
	repository.NewDevice(db)
}
