package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/markbates/pkger"

	"github.com/scinna/server/middleware"
	"github.com/scinna/server/routes"
	"github.com/scinna/server/services"
)

// RunServer starts the web api server
func RunServer(prv *services.Provider) {

	r := mux.NewRouter().StrictSlash(false)

	// the react fontend app
	r.HandleFunc("/", routes.IndexRoute(prv))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	r.HandleFunc("/config", middleware.CombineMiddlewaresCT(prv, routes.GetConfigRoute(prv))).Methods("GET")

	authRoutes := r.PathPrefix("/auth").Subrouter().StrictSlash(false)
	authRoutes.HandleFunc("/login", middleware.CombineMiddlewaresCT(prv, routes.LoginRoute(prv))).Methods("POST")
	authRoutes.HandleFunc("/register", middleware.CombineMiddlewaresCT(prv, routes.IsRegisterAvailableRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/register", middleware.CombineMiddlewaresCT(prv, routes.RegisterRoute(prv))).Methods("POST")
	authRoutes.HandleFunc("/register/{VALIDATION_TOKEN}", routes.ValidateUserRoute(prv))
	authRoutes.HandleFunc("/tokens", middleware.CombineMiddlewaresCT(prv, routes.GetTokensRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/tokens/{TOKEN_ID}", middleware.CombineMiddlewaresCT(prv, routes.RevokeTokenRoute(prv))).Methods("DELETE")

	picturesRoutes := r.PathPrefix("/pictures").Subrouter().StrictSlash(false)
	// For some reason using / forces the trailing slash in the url.. So blanking it out...
	picturesRoutes.HandleFunc("", middleware.CombineMiddlewaresCT(prv, routes.UploadPictureRoute(prv))).Methods("POST")
	picturesRoutes.HandleFunc("/{URL_ID}", middleware.CombineMiddlewaresCT(prv, routes.DeletePictureRoute(prv))).Methods("DELETE")
	picturesRoutes.HandleFunc("/{URL_ID}", middleware.CombineMiddlewaresCT(prv, routes.PictureInfoRoute(prv))).Methods("GET")

	usersRoutes := r.PathPrefix("/users").Subrouter().StrictSlash(false)
	usersRoutes.HandleFunc("/me", middleware.CombineMiddlewaresCT(prv, routes.UpdateMyInfosRoute(prv))).Methods("PUT")
	usersRoutes.HandleFunc("/{username}", middleware.CombineMiddlewaresCT(prv, routes.UserInfoRoute(prv))).Methods("GET")

	adminRoutes := r.PathPrefix("/admin").Subrouter().StrictSlash(false)
	adminRoutes.HandleFunc("/invite", middleware.CombineMiddlewaresCT(prv, routes.GenerateInviteRoute(prv))).Methods("POST")

	// Default route is for picture laoding
	r.HandleFunc("/{pict}", middleware.CombineMiddlewares(prv, routes.RawPictureRoute(prv), false)).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + prv.Config.WebPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

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
		Addr:         "127.0.0.1:" + strconv.Itoa(*port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
