package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

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
	authRoutes.HandleFunc("/logout", middleware.CombineMiddlewaresCT(prv, routes.LogoutRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/token", middleware.CombineMiddlewaresCT(prv, routes.CheckTokenRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/register", middleware.CombineMiddlewaresCT(prv, routes.IsRegisterAvailableRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/register", middleware.CombineMiddlewaresCT(prv, routes.RegisterRoute(prv))).Methods("POST")
	authRoutes.HandleFunc("/register/{VALIDATION_TOKEN}", routes.ValidateUserRoute(prv))
	authRoutes.HandleFunc("/tokens", middleware.CombineMiddlewaresCT(prv, routes.GetTokensRoute(prv))).Methods("GET")
	authRoutes.HandleFunc("/tokens/{TOKEN_ID}", middleware.CombineMiddlewaresCT(prv, routes.RevokeTokenRoute(prv))).Methods("DELETE")

	mediasRoutes := r.PathPrefix("/medias").Subrouter().StrictSlash(false)
	// For some reason using / forces the trailing slash in the url.. So blanking it out...
	mediasRoutes.HandleFunc("", middleware.CombineMiddlewaresCT(prv, routes.UploadMediaRoute(prv))).Methods("POST")
	mediasRoutes.HandleFunc("/{URL_ID}", middleware.CombineMiddlewaresCT(prv, routes.DeleteMediaRoute(prv))).Methods("DELETE")
	mediasRoutes.HandleFunc("/{URL_ID}", middleware.CombineMiddlewaresCT(prv, routes.MediaInfoRoute(prv))).Methods("GET")

	// Folders routes
	r.PathPrefix("/folders").Handler(
		http.StripPrefix("/folders", middleware.CombineMiddlewaresCT(prv, routes.GetFolderContentRoute(prv))),
	).Methods("GET")

	r.PathPrefix("/folders").Handler(
		http.StripPrefix("/folders", middleware.CombineMiddlewaresCT(prv, routes.CreateFolderRoute(prv))),
	).Methods("POST")

	foldersRoutes := r.PathPrefix("/folders").Subrouter().StrictSlash(false)
	foldersRoutes.HandleFunc("/{FOLDER_ID}/{NEW_NAME}", middleware.CombineMiddlewaresCT(prv, routes.RenameFolderRoute(prv))).Methods("UPDATE")
	foldersRoutes.HandleFunc("/{FOLDER_ID}", middleware.CombineMiddlewaresCT(prv, routes.MoveFolderRoute(prv))).Methods("PUT")              // Move to root
	foldersRoutes.HandleFunc("/{FOLDER_ID}/{NEW_PARENT}", middleware.CombineMiddlewaresCT(prv, routes.MoveFolderRoute(prv))).Methods("PUT") // Move to a folder
	foldersRoutes.HandleFunc("/{FOLDER_ID}", middleware.CombineMiddlewaresCT(prv, routes.DeleteFolderRoute(prv))).Methods("DELETE")

	usersRoutes := r.PathPrefix("/users").Subrouter().StrictSlash(false)
	usersRoutes.HandleFunc("/me", middleware.CombineMiddlewaresCT(prv, routes.UpdateMyInfosRoute(prv))).Methods("PUT")
	usersRoutes.HandleFunc("/{username}", middleware.CombineMiddlewaresCT(prv, routes.UserInfoRoute(prv))).Methods("GET")

	adminRoutes := r.PathPrefix("/admin").Subrouter().StrictSlash(false)
	adminRoutes.HandleFunc("/invite", middleware.CombineMiddlewaresCT(prv, routes.GenerateInviteRoute(prv))).Methods("POST")

	// Default route is for media laoding
	r.HandleFunc("/{media}", middleware.CombineMiddlewares(prv, routes.RawMediaRoute(prv), false)).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + prv.Config.WebPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
