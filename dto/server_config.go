package dto

import "github.com/scinna/server/services"

type ServerConfig struct {
	RegistrationAllowed bool
	Validation          string
	WebURL              string
}

func NewServerConfig(prv *services.Provider) ServerConfig {
	return ServerConfig{
		RegistrationAllowed: prv.Config.Registration.Allowed,
		Validation: prv.Config.Registration.Validation,
		WebURL: prv.Config.WebURL,
	}
}
