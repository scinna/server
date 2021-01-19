package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/scinna/server/config"
	"github.com/scinna/server/cron"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/log"
	"github.com/scinna/server/routes"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
)

const (
	SCINNA_AUTHOR  = "Scinna Team"
	SCINNA_VERSION = "0.1"
	SCINNA_PATCH   = "0"
)

func main() {
	err := start()
	if err != nil {
		fmt.Println(err)
	}
}

func start() error {
	fmt.Printf("Scinna [v%v.%v] by %v\n", SCINNA_VERSION, SCINNA_PATCH, SCINNA_AUTHOR)

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	prv, err := services.NewProvider(cfg)
	if err != nil {
		return err
	}

	err = utils.CheckVersion(prv, SCINNA_VERSION)
	if err != nil {
		return err
	}

	SetupCloseHandler(prv)

	go func() {
		for {
			cron.ClearOldAccounts(prv)
			time.Sleep(1 * time.Hour)
		}
	}()

	if !prv.Config.Registration.Allowed {
		code, err := dal.GenerateInviteIfNeeded(prv)
		if err != nil {
			return err
		}

		if code != "NONE" {
			log.InfoAlwaysShown("Your server is invite only and there are no users registered.")
			log.InfoAlwaysShown("Please use this code to create your admin account: " + code)
		}
	}

	router := mux.NewRouter()

	routes.WebApp(prv, router)
	routes.Authentication(prv, router.PathPrefix("/auth").Subrouter())
	routes.Accounts(prv, router.PathPrefix("/account").Subrouter())

	// Last one (Matching the media_id)
	routes.Medias(prv, router.PathPrefix("/").Subrouter())

	srv := &http.Server{
		Handler:           router,
		Addr:              cfg.ListeningAddr,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	srv.ListenAndServe()

	return nil
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
