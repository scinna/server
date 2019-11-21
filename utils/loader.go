package utils

import (
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func LoadDatabase() *sqlx.DB {
	dsn, exists := os.LookupEnv("POSTGRES_DSN")
	if !exists {
		panic("No database url found! (POSTGRES_DSN)")
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}