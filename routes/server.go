package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"net/http"
)

func Server(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/infos", configRoute(prv))

	r.HandleFunc("/logo", logoWide(prv))
	r.HandleFunc("/logo/wide", logoWide(prv))
	r.HandleFunc("/logo/small", logoSmall(prv))
}

func configRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _, _ := middlewares.GetUserFromRequest(prv, r)
		isAdmin := false
		if user != nil {
			isAdmin = user.IsAdmin
		}

		w.Header().Set("Content-Type", "application/json")

		bytes, err := json.Marshal(dto.NewServerConfig(prv, isAdmin))
		if serrors.WriteError(w, r, err) {
			return
		}

		_, _ = w.Write(bytes)
	}
}

func logoWide(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(prv.Config.CustomLogoWide) > 0 {
			http.ServeFile(w, r, prv.Config.CustomLogoWide)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(prv.EmbeddedAssets.LogoWide)
	}
}

func logoSmall(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(prv.Config.CustomLogoSmall) > 0 {
			http.ServeFile(w, r, prv.Config.CustomLogoSmall)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(prv.EmbeddedAssets.LogoSmall)
	}
}