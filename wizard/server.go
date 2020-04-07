package wizard

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"
	"github.com/scinna/server/routes"
)

// RunSetup spin up a http server letting the admin do the first time setup of the server. This keeps the routes separated and so will never be called again as soon as the setup is done
func RunSetup(port *int) {

	var srv http.Server

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/test/db", routes.TestDatabaseConfigRoute)
	r.HandleFunc("/test/smtp", routes.TestSMTPConfigRoute)
	r.HandleFunc("/scinna", routes.ScinnaConfigRoute)
	r.HandleFunc("/user", routes.CreateAdminRoute)
	r.PathPrefix("/").Handler(http.FileServer(pkger.Dir("/frontend/setup/build"))) // @TODO: Embed files in the executable through something like markbates/pkger

	srv = http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + strconv.Itoa(*port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
