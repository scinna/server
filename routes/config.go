package routes

import (
	"encoding/json"
	"net/http"

	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

type configExposed struct {
	HasSMTP      bool   `json:"EmailAvailable"`
	Registration string `json:"Registration"`
}

// GetConfigRoute return the configuration of the server useful for apps
func GetConfigRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := configExposed{
			HasSMTP:      prv.Config.Mail.Enabled,
			Registration: prv.Config.RegistrationAllowed,
		}

		cfgJSON, err := json.Marshal(cfg)
		if err != nil {
			serrors.WriteError(w, err)
			return
		}

		w.Write(cfgJSON)
	}
}
