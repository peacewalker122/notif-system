package config

import (
	"fmt"
)

type Config struct {
	DBConfig *DB
	APP      *APP
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string

	PreparedStatements []string
}

type APP struct {
	Host string
}

func (d *DB) URL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
}
