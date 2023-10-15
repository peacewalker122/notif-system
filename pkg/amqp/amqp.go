package amqp

import (
	"context"
	"log"

	"notifsys/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQP struct {
	Conn *amqp.Connection
	done chan struct{}
}

func New(done chan struct{}) *AMQP {
	amqpcfg := config.Get().AMQP
	conn, err := amqp.Dial(amqpcfg.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &AMQP{
		Conn: conn,
		done: done,
	}
}

// Consume consumes messages from the specified queue.
//
// ctx: The context.Context to use for the operation.
// queue: The name of the queue to consume from.
// Returns: A channel of amqp.Delivery and an error.
func (a *AMQP) Consume(ctx context.Context, queue string) (<-chan amqp.Delivery, func(), error) {
	channel, err := a.Conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	// defer channel.Close()

	msgs, err := channel.Consume(
		queue,
		"gonotifsys", // Use a unique consumer tag
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return msgs, func() {
		channel.Close()
	}, nil
}

func (a *AMQP) Produce(ctx context.Context, queue string, data []byte) error {
	channel, err := a.Conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.PublishWithContext(
		context.Background(),
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
