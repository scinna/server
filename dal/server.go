package dal

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/scinna/server/services"
)

// GetDbVersion returns the database version. It returns -1 if the DB is not setup.
func GetDbVersion(prv *services.Provider) (int, error) {
	var dbv int
	err := prv.Db.Get(&dbv, "SELECT DBVERSION FROM DBINFO")

	if err != nil {
		// If there is an error and it is something like "Relation dbinfo does not exists" or "No rows" this mean that the
		// server has never been initialized, thus we need to let the user set it up
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "undefined_table" {
				return -1, nil
			}
		}

		if err == sql.ErrNoRows {
			return -1, nil
		}

		return 0, err
	}
	return dbv, nil
}
