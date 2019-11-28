package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/oxodao/scinna/configuration"
	"github.com/oxodao/scinna/middleware"
	"github.com/oxodao/scinna/routes"
	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/utils"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Scinna Server - V1")

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using currently exported vars")
	}
	fmt.Println("- Env var loaded")

	cfg, err := configuration.Load()
	if err != nil {
		panic(err)
	}

	utils.GenerateDefaultPicture()

	argonParams := &services.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	db := utils.LoadDatabase(cfg.PostgresDSN)
	defer db.Close()
	fmt.Println("- Connected to database")

	prv := services.New(cfg, db, utils.LoadMail(), argonParams)

	r := mux.NewRouter().StrictSlash(false)

	// the react fontend app
	r.HandleFunc("/", routes.IndexRoute(prv))

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
	usersRoutes.HandleFunc("/me/pictures", middleware.CombineMiddlewaresCT(prv, routes.MyPicturesRoute(prv))).Methods("GET")
	usersRoutes.HandleFunc("/me", middleware.CombineMiddlewaresCT(prv, routes.MyInfosRoute(prv))).Methods("GET")
	usersRoutes.HandleFunc("/me", middleware.CombineMiddlewaresCT(prv, routes.UpdateMyInfosRoute(prv))).Methods("PUT")
	usersRoutes.HandleFunc("/{username}/pictures", middleware.CombineMiddlewaresCT(prv, routes.UserPicturesRoute(prv))).Methods("GET")

	adminRoutes := r.PathPrefix("/admin").Subrouter().StrictSlash(false)
	adminRoutes.HandleFunc("/invite", middleware.CombineMiddlewaresCT(prv, routes.GenerateInviteRoute(prv))).Methods("POST")

	// Default route is for picture laoding
	r.HandleFunc("/{pict}", middleware.CombineMiddlewares(prv, routes.RawPictureRoute(prv), false)).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + cfg.WebPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
