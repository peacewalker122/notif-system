package amqp

import (
	"log"

	"notifsys/internal/config"

	"github.com/streadway/amqp"
	"github.com/uptrace/bun"
)

type AMQP struct {
	db   *bun.DB
	conn *amqp.Connection
	done chan struct{}
}

func New(db *bun.DB, done chan struct{}) *AMQP {
	amqpcfg := config.Get().AMQP
	conn, err := amqp.Dial(amqpcfg.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &AMQP{
		db:    db,
		conn:  conn,
		done:  done,
	}
}

func (a *AMQP) Consume(queue string) (<-chan amqp.Delivery, error) {
	channel, err := a.conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		queue,
		"go-notif",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgs, err
}
