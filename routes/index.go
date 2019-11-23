package routes

import (
	"github.com/oxodao/scinna/services"
	"net/http"
)

// IndexRoute is the index endpoint, the one displaying the react webapp
func IndexRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This will be the react app for Scinna picture server"))
	}
}
