package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	headerIPField, exists := os.LookupEnv("HEADER_IP_FIELD")
	if !exists {
		fmt.Println("The header for the IP field is not set (HEADER_IP_FIELD). If you are using a reverse-proxy please be sure to set it according to its configuration.\nTo disable this message, add the environment variable with an empty value.")
	}

	registrationAllowed, exists := os.LookupEnv("REGISTRATION_ALLOWED")
	var registrationAllowedBool bool
	if !exists {
		fmt.Println("Registration is allowed by default. You can't hide this message or turn it off by filling the \"REGISTRATION_ALLOWED\" environment variable.")
	} else {
		registrationAllowedBool, err := strconv.ParseBool(registrationAllowed)
		if err != nil {
			panic("Can't parse REGISTRATION_ALLOWED. It should be either true or false")
		}
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

	prv := services.New(db, argonParams, picturepath, headerIPField, registrationAllowedBool)

	r := mux.NewRouter().StrictSlash(false)

	// the react fontend app
	r.HandleFunc("/", routes.IndexRoute(prv))

	authRoutes := r.PathPrefix("/auth").Subrouter().StrictSlash(false)
	authRoutes.Use(middleware.ContentTypeMiddleware)
	authRoutes.HandleFunc("/login", routes.LoginRoute(prv)).Methods("POST")
	authRoutes.HandleFunc("/register", routes.IsRegisterAvailableRoute(prv)).Methods("GET")
	authRoutes.HandleFunc("/register", routes.RegisterRoute(prv)).Methods("POST")
	authRoutes.HandleFunc("/register/{VALIDATION_TOKEN}", routes.ValidationRoute(prv)).Methods("POST")

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
