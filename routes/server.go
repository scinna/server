package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"net/http"
)

func Server(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/infos", configRoute(prv))
	r.HandleFunc("/logo", logo(prv))
}

func configRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bytes, err := json.Marshal(dto.NewServerConfig(prv))
		if serrors.WriteError(w, r, err) {
			return
		}

		_, _ = w.Write(bytes)
	}
}
func logo (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(prv.Config.CustomLogo) > 0 {
			http.ServeFile(w, r, prv.Config.CustomLogo)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(prv.LogoFile)
	}
}