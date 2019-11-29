package utils

import (
	"github.com/jmoiron/sqlx"
)

// LoadDatabase create the sqlx DB instance
func LoadDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
