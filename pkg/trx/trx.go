package trx

import (
	"context"
	"fmt"
	"runtime"

	"github.com/uptrace/bun"
)

type Manager struct {
	db *bun.DB
}

func New(db *bun.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (s *Manager) Run(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Println("[BEGIN]")
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			runtime.Stack(buf, false)
			fmt.Println(string(buf))
			err = tx.Rollback()
			fmt.Println("[ROLLBACK]")
			return
		}
		if err != nil {
			err = tx.Rollback()
			fmt.Println("[ROLLBACK]")
			return
		}
		err = tx.Commit()
		fmt.Println("[COMMIT]")
		return
	}()

	err = fn(ctx)

	return
}
