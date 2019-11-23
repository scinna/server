package utils

import (
	"os"

	"github.com/jmoiron/sqlx"
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

/**
 *  This method checks if the picture displayed for
 *  non existing image exists.
 *  If it doesn't, it generates a basic one.
 **/
func GenerateDefaultPicture() {
	// @TODO: Generate a default picture if it doesn't exists
	pict, err := os.Open("not_found.png")
	if err != nil {
		panic(err)
	}
	pict.Close()
}
