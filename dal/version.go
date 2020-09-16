package dal

import "github.com/scinna/server/services"

func FetchVersion(prv *services.Provider) (string, error) {
	row := prv.DB.QueryRowx("SELECT VERSION FROM DBVERSION LIMIT 1")
	var version string
	if row.Err() != nil {
		return "", row.Err()
	}

	row.Scan(&version)

	return version, nil
}
