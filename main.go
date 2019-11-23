package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	port, exists := os.LookupEnv("WEB_PORT")
	if !exists {
		panic("No listening port found! (WEB_PORT)")
	}

	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		panic("No JWT secret found! (JWT_SECRET)\nGenerate one with this command => openssl rand -base64 172 | tr -d '\\n'")
	}
	jwtSecretDecoded, err := base64.StdEncoding.DecodeString(jwtSecret)
	if err != nil {
		panic("BAD JWT SECRET!")
	}

	utils.GenerateDefaultPicture()

	picturepath, exists := os.LookupEnv("PICTURE_PATH")
	if !exists {
		panic("No picture folder found! (PICTURE_PATH)")
	}

	argonParams := &services.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	db := utils.LoadDatabase()
	defer db.Close()
	fmt.Println("- Connected to database")

	prv := services.New(db, argonParams, jwtSecretDecoded, picturepath)

	r := mux.NewRouter().StrictSlash(false)

	// the react fontend app
	r.HandleFunc("/", routes.IndexRoute(prv))

	authRoutes := r.PathPrefix("/auth").Subrouter().StrictSlash(false)
	authRoutes.Use(middleware.ContentTypeMiddleware)
	authRoutes.HandleFunc("/login", routes.LoginRoute(prv)).Methods("POST")     // Login route to get a JWT token
	authRoutes.HandleFunc("/refresh", routes.RefreshRoute(prv)).Methods("POST") // Refresh route to refresh the JWT token

	picturesRoutes := r.PathPrefix("/pictures").Subrouter().StrictSlash(false)
	picturesRoutes.Use(middleware.ContentTypeMiddleware)
	// For some reason using / forces the trailing slash in the url.. So blanking it out...
	picturesRoutes.HandleFunc("", routes.UploadPictureRoute(prv)).Methods("POST")
	picturesRoutes.HandleFunc("/{URL_ID}", routes.DeletePictureRoute(prv)).Methods("DELETE")
	picturesRoutes.HandleFunc("/{URL_ID}", routes.PictureInfoRoute(prv)).Methods("GET")

	usersRoutes := r.PathPrefix("/users").Subrouter().StrictSlash(false)
	usersRoutes.Use(middleware.ContentTypeMiddleware)
	usersRoutes.HandleFunc("/me/pictures", routes.MyPicturesRoute(prv)).Methods("GET")
	usersRoutes.HandleFunc("/me", routes.MyInfosRoute(prv)).Methods("GET")
	usersRoutes.HandleFunc("/me", routes.UpdateMyInfosRoute(prv)).Methods("PUT")
	usersRoutes.HandleFunc("/{username}/pictures", routes.UserPicturesRoute(prv)).Methods("GET")

	// Default route is for picture laoding
	r.HandleFunc("/{pict}", routes.RawPictureRoute(prv)).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
