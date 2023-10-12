package config

import (
	"os"
	"sync"
)

var (
	once = new(sync.Once)
	cfg  = new(Config)
)

func Get() *Config {
	once.Do(func() {
		cfg = &Config{
			DBConfig: &DB{
				Host:     os.Getenv("DBHOST"),
				Port:     os.Getenv("DBPORT"),
				User:     os.Getenv("DBUSER"),
				Password: os.Getenv("DBPASSWORD"),
				DBName:   os.Getenv("DBNAME"),
				SSLMode:  os.Getenv("DBSSLMODE"),
				PreparedStatements: []string{
					`PREPARE insert_room as INSERT INTO rooms (id, name) VALUES ($1, $2);`,
				},
			},
			APP: &APP{
				Host: os.Getenv("APPHOST"),
			},
		}
	})
	return cfg
}
