package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/services"
)

func CheckVersion(prv *services.Provider, scinnaVersion string) error {
	version, err := dal.FetchVersion(prv)
	if err != nil {
		errFound := false
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "42P01" { // Relation does not exists == database is not created
				errFound = true
				log.Fatal("You should initialize the database with the given script")
			}
		}

		if !errFound {
			return err
		}
	}

	if version != scinnaVersion {
		return errors.New(fmt.Sprintf(`Your database is not up to date. Please execute migrations
		You should be on v.%v.x. You are currently running %v.`, version, scinnaVersion))
	}

	return nil
}
