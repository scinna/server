package routes

import (
	"net/http"

	"github.com/scinna/server/dal"
	"github.com/scinna/server/services"
)

// DatabaseVersion is used for two things: Checking if the database is initialized, and checking if the server just got an update in order to execute the migrations
const DatabaseVersion int = 1

// IndexRoute is the index endpoint, the one displaying the react webapp
func IndexRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbv, err := dal.GetDbVersion(prv)

		if err != nil {
			// @TODO: Show an error to the user
			w.WriteHeader(500)
			return
		}

		if dbv == -1 {
			http.Redirect(w, r, "setup", http.StatusSeeOther)
			return
		}

		if dbv != DatabaseVersion {
			http.Redirect(w, r, "migrate", http.StatusSeeOther)
			return
		}

		w.Write([]byte("This will be the react app for Scinna picture server"))
	}
}

// SetupRoute is a route letting the user setup the server. It should only be ran once
func SetupRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If the server is already setup, you must not be able to get into this page
		dbv, err := dal.GetDbVersion(prv)
		if dbv > 0 || err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		w.Write([]byte("Lets setup"))
	}
}

// MigrateRoute is a route letting the user update or rollback the server once he change versions
func MigrateRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Lets migrate"))

		// @TODO: Migrations / Rollback
		// Download the missing rollback files maybe ?
	}
}
