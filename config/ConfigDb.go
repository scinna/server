package config

import "errors"

/** DB represents the configuration for the database **/
type DB struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

func (db *DB) Validate() error {
	var err []error
	if len(db.Hostname) == 0 {
		err = append(err, errors.New("Database.Hostname can't be empty"))
	}

	if db.Port == 0 {
		err = append(err, errors.New("Database.Port can't be empty"))
	}

	if len(db.Username) == 0  {
		err = append(err, errors.New("Database.Username can't be empty"))
	}

	if len(db.Password) == 0  {
		err = append(err, errors.New("Database.Password can't be empty"))
	}

	if len(db.Database) == 0  {
		err = append(err, errors.New("Database.Database can't be empty"))
	}

	return combineErrors(err...)
}
