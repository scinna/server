package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
)

func Accounts(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.LoggedInMiddleware(prv))
	r.Use(middlewares.Json)

	r.HandleFunc("/", fetchAccountInfos(prv)).Methods(http.MethodGet)
}

func fetchAccountInfos(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		if user == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		u, _ := json.Marshal(user)
		w.Write(u)
	}
}
