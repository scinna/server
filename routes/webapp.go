package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

func WebApp(prv *services.Provider, r *mux.Router) {
	// /app will show the frontend app
	// We can't use / since it will be used to serve pictures and it would conflict with assets serving
	// A solution would be to move al assets to a sub-directory and configure the frontend to let access on /assets for example
	r.HandleFunc("/app", homeRoute(prv))
	r.HandleFunc("/infos", configRoute(prv))
}

func homeRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to scinna webapp"))
		// @TODO use pkger to serve the webapp
	}
}

func configRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(dto.NewServerConfig(prv))
		if serrors.WriteError(w, err) {
			return
		}

		w.Write(bytes)
	}
}
