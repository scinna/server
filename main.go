package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/scinna/server/config"
	"github.com/scinna/server/cron"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/log"
	"github.com/scinna/server/routes"
	"github.com/scinna/server/services"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	SCINNA_AUTHOR = "Scinna Team"
	SCINNA_VERSION = "0.1"
	SCINNA_PATCH   = "0"
)

func main() {
	fmt.Printf("Scinna [v%v.%v] by %v\n", SCINNA_VERSION, SCINNA_PATCH, SCINNA_AUTHOR)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	prv, err := services.NewProvider(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	version, err := dal.FetchVersion(prv)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if version != SCINNA_VERSION {
		log.Fatal("Your database is not up to date. Please execute migrations")
		log.Fatal("You should be on v." + version + ".x")
		return
	}

	SetupCloseHandler(prv)

	go func() {
		for {
			cron.ClearOldAccounts(prv)
			time.Sleep(1 * time.Hour)
		}
	}()

	router := mux.NewRouter()

	routes.WebApp(prv, router)
	routes.Authentication(prv, router.PathPrefix("/auth").Subrouter())
	routes.Accounts(prv, router.PathPrefix("/account").Subrouter())

	// Last one (Matching the media_id)
	routes.Medias(prv, router.PathPrefix("/").Subrouter())

	headers := handlers.AllowedHeaders([]string {
		"Authorization",
		"X-Requested-With",
		"X-Real-IP",
		"Content-Type",
	})

	methods := handlers.AllowedMethods([]string{
		"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE", "PATCH",
	})

	origins := handlers.AllowedOrigins([]string{
		"*",
	})

	srv := &http.Server{
		Handler: handlers.CORS(headers, methods, origins)(router),
		Addr: "0.0.0.0:" + strconv.Itoa(cfg.WebPort),
		WriteTimeout: 15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	srv.ListenAndServe()
}

func SetupCloseHandler(prv *services.Provider) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Shutting down scinna")
		prv.Shutdown()
		os.Exit(0)
	}()
}