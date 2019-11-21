package routes

import (
	"net/http"
	"github.com/oxodao/scinna/services"
)

func IndexRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This will be the react app for Scinna picture server"))
	}
}
