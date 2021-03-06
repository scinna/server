package config

import (
	"errors"
	"strings"
)

type Registration struct {
	HasUser    bool
	Allowed    bool
	Validation string
}

func (r Registration) Validate() error {
	if strings.ToLower(r.Validation) != "open" && strings.ToLower(r.Validation) != "email" && strings.ToLower(r.Validation) != "admin" {
		return errors.New("Registration.Validation must be either \"open\", \"email\" or \"admin\"")
	}

	r.Validation = strings.ToLower(r.Validation)

	return nil
}
