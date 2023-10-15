package amqp

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"notifsys/internal/dto"

	"github.com/joho/godotenv"
)

func TestAMQP(t *testing.T) {
	godotenv.Load("../../.env.local")

	test := New(make(chan struct{}))

	data, err := json.Marshal(dto.NotifRequest{
		UserID: []int{907551073050263553},
		Message: map[string]string{
			"title": "title",
			"body":  "body11",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		test.Produce(context.Background(), "notif", data)
	}
}

func TestConsume(t *testing.T) {
	godotenv.Load("../../.env.local")
	test := New(make(chan struct{}))

	del, done, err := test.Consume(context.Background(), "notif")
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	go func() {
		for d := range del {
			go t.Log(string(d.Body))
		}
	}()

	select {
	case <-time.After(7 * time.Second):
		t.Log("Success")
	}
}
