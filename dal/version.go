package dal

import "github.com/scinna/server/services"

func FetchVersion(prv *services.Provider) (string, error) {
	row := prv.DB.QueryRow("SELECT VERSION FROM dbversion")
	if row.Err() != nil {
		return "", row.Err()
	}

	var vers string
	err := row.Scan(&vers)

	return vers, err
}
