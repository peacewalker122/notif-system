package cron

import (
	"context"

	"notifsys/internal/factory"

	"github.com/uptrace/bun"
)

func NewCron(f *factory.Factory) *Cron {
	return &Cron{
		db:      f.DB,
		service: NewService(f),
		done:    f.Done,
	}
}

type Cron struct {
	db      *bun.DB
	done    chan struct{}
	service Service
}

func (c *Cron) Run(ctx context.Context) {
	go func() {
		c.notifHandler(ctx)
	}()

	// job.RunAsync()
	<-c.done
}

func (c *Cron) notifHandler(ctx context.Context) {
	c.service.Notification(ctx)
}
