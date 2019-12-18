package routes

import (
	"github.com/oxodao/scinna/services"
	"net/http"
)

// DatabaseVersion is used for two things: Checking if the database is initialized, and checking if the server just got an update in order to execute the migrations
const DatabaseVersion int = 1

// IndexRoute is the index endpoint, the one displaying the react webapp
func IndexRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This will be the react app for Scinna picture server"))
	}
}
