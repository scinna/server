package config

import (
	"errors"
	"strings"
)

/** SMTP represents the configuration for the database **/
type SMTP struct {
	Enabled        bool
	ConnectionType string

	Hostname string
	Port     int
	Username string
	Password string
	Sender   string
}

func (smtp SMTP) Validate() error {
	if smtp.Enabled {
		var err []error

		if strings.ToLower(smtp.ConnectionType) != "starttls" && strings.ToLower(smtp.ConnectionType) != "plain" {
			err = append(err, errors.New("Mail.ConnectionType must be either \"STARTTLS\" or \"PLAIN\""))
		}

		smtp.ConnectionType = strings.ToLower(smtp.ConnectionType)

		if len(smtp.Hostname) == 0  {
			err = append(err, errors.New("Mail.Hostname can't be empty"))
		}

		if smtp.Port == 0  {
			err = append(err, errors.New("Mail.Port can't be empty"))
		}

		if len(smtp.Sender) == 0  {
			err = append(err, errors.New("Mail.Sender can't be empty"))
		}

		return combineErrors(err...)
	}

	return nil
}
