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
	// A solution would be to move all assets to a sub-directory and configure the frontend to let access on /assets for example
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app/", http.StatusPermanentRedirect)
	})
	r.PathPrefix("/app").Handler(http.StripPrefix("/app/", http.FileServer(http.FS(*prv.Webapp))))
	r.HandleFunc("/api/infos", configRoute(prv))
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
