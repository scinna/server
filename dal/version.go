package dal

import (
	"github.com/jmoiron/sqlx"
)

func FetchVersion(db *sqlx.DB) (string, error) {
	row := db.QueryRow("SELECT VERSION FROM dbversion")
	if row.Err() != nil {
		return "", row.Err()
	}

	var vers string
	err := row.Scan(&vers)

	return vers, err
}
