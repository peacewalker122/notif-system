package test

import (
	"context"
	"log"
	"testing"
	"time"

	"notifsys/internal/app/cron"
	"notifsys/internal/config"
	"notifsys/internal/factory"
	"notifsys/pkg/db"
	"notifsys/pkg/fcm"
	"notifsys/pkg/tracer"

	"github.com/joho/godotenv"
	"github.com/uptrace/uptrace-go/uptrace"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	godotenv.Load("../../../../.env.local")
	donechan := make(chan struct{})

	database := db.New()

	err := fcm.New("../../../../private/fcm.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	f := factory.New(database.DB, donechan)
	tracer.New()
	s := cron.NewService(f)

	cfg := config.Get().APP

	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(cfg.Name),
	)
	defer uptrace.Shutdown(ctx)

	go func() {
		select {
		case <-time.After(5 * time.Second):
			close(donechan)
		}
	}()

	t.Run("notification", func(t *testing.T) {
		t.Parallel()
		s.Notification(ctx)
	})
}
