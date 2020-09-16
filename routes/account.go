package routes

import (
	"github.com/gorilla/mux"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
	"net/http"
)

func Accounts(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.LoggedInMiddleware(prv))

	r.HandleFunc("/", fetchAccountInfos(prv)).Methods(http.MethodGet)
}

func fetchAccountInfos(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		w.Write([]byte("Hehehe " + user.Name))
	}
}

