package dto

import (
	"fmt"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

type ServerConfig struct {
	RegistrationAllowed bool
	Validation          string
	WebURL              string
	CustomBranding      string `json:",omitempty"`
	ScinnaVersion       string `json:",omitempty"`
}

func NewServerConfig(prv *services.Provider, isAdmin bool) ServerConfig {
	cfg := ServerConfig{
		RegistrationAllowed: prv.Config.Registration.Allowed,
		Validation:          prv.Config.Registration.Validation,
		WebURL:              prv.Config.WebURL,
		CustomBranding:      prv.Config.CustomBranding,
	}

	if isAdmin {
		cfg.ScinnaVersion = fmt.Sprintf("%v.%v", utils.SCINNA_VERSION, utils.SCINNA_PATCH)
	}

	return cfg
}
