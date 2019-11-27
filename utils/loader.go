package utils

import (
	"os"

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

// GenerateDefaultPicture method checks if the picture displayed for non existing image exists. If it doesn't, it generates a basic one
func GenerateDefaultPicture() {
	pict, err := os.Open("not_found.png")
	if err != nil {
		panic(err)
	}
	pict.Close()
}
