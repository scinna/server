package services

import (
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/matoous/go-nanoid"
)

type Provider struct {
	Db     *sqlx.DB
}

func (prv *Provider) GenerateUID() (string, error) {
	alphabet, exists := os.LookupEnv("ID_ALPHABET")
	if !exists {
		alphabet = "0123456789_abcdefghijklmnopqrstuvwxyz-ABCEFGHIJKLMNOPQRSTUVWXYZ"
	}
	
	lengthStr, exists := os.LookupEnv("ID_LENGTH")
	length := 10
	if exists {
		length, _ = strconv.Atoi(lengthStr)
	}

	return gonanoid.Generate(alphabet, length)
}

// Temporary empty, will contain Postgres connection
func New(db *sqlx.DB) *Provider {
	return &Provider{
		Db:     db,
	}
}
