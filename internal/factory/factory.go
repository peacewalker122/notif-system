package factory

import (
	"notifsys/internal/repository"
	"notifsys/internal/repository/interfaces"
	"notifsys/pkg/amqp"
	"notifsys/pkg/fcm"

	"github.com/uptrace/bun"
)

type Factory struct {
	DB   *bun.DB
	AMQP *amqp.AMQP

	interfaces.Device
	interfaces.User
	interfaces.FCM

	Done chan struct{}
}

func New(db *bun.DB, done chan struct{}) *Factory {
	f := new(Factory)
	f.Done = done
	f.Setup(db)

	return f
}

func (f *Factory) Setup(db *bun.DB) {
	f.DB = db
	f.AMQP = amqp.New(f.Done)
	f.Device = repository.NewDevice(f.DB)
	f.User = repository.NewUser(f.DB)
	f.FCM = repository.New(fcm.APP)
}
