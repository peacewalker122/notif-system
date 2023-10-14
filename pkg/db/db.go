package db

import (
	"database/sql"
	"time"

	"notifsys/internal/config"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bunotel"
)

type DB struct {
	*bun.DB
}

func New() *DB {
	dbCfg := config.Get().DBConfig

	sqldb, err := sql.Open("postgres", dbCfg.URL())
	if err != nil {
		panic(err)
	}
	sqldb.SetConnMaxIdleTime(5 * time.Minute)
	sqldb.SetMaxOpenConns(10)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetConnMaxLifetime(5 * time.Minute)

	db := bun.NewDB(sqldb, pgdialect.New())

	// db.AddQueryHook(bundebug.NewQueryHook(
	// 	bundebug.WithVerbose(true),
	// 	bundebug.FromEnv("2"),
	// ))
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("app"),
	))

	return &DB{
		DB: db,
	}
}

func (d *DB) Close() {
	d.DB.Close()
}
