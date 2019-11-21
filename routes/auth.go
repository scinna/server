package routes

import (
	"github.com/oxodao/scinna/services"
	"net/http"
)

func LoginRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("LoginRoute - To be implemented"))
	}
}

func RefreshRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("RefreshRoute - To be implemented"))
	}
}