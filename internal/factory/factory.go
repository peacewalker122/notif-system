package factory

import (
	"notifsys/internal/repository"
	"notifsys/internal/repository/interfaces"
	"notifsys/pkg/fcm"

	"github.com/uptrace/bun"
)

type Factory struct {
	DB *bun.DB

	interfaces.Device
	interfaces.User
	interfaces.FCM
}

func New(db *bun.DB, done chan struct{}) *Factory {
	f := new(Factory)
	f.Setup(db)

	return f
}

func (f *Factory) Setup(db *bun.DB) {
	f.DB = db
	f.Device = repository.NewDevice(f.DB)
	f.User = repository.NewUser(f.DB)
	f.FCM = repository.New(fcm.APP)
}
