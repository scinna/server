package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/scinna/server/translations"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/scinna/server/config"
	"github.com/scinna/server/cron"
	"github.com/scinna/server/fixtures"
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

//go:embed frontend/dist
var frontend embed.FS

func main() {
	err := start()
	if err != nil {
		fmt.Println(err)
	}
}

func start() error {
	fmt.Printf("Scinna [v%v.%v] by %v\n", SCINNA_VERSION, SCINNA_PATCH, SCINNA_AUTHOR)

	translations.Initialize()

	generateDb := flag.Bool("generate-db", false, "Generate the default database")
	forceGenerateDb := flag.Bool("auto-yes", false, "Automatically answer yes to the dropping of the old tables (CAUTION!)")

	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	prv, err := services.NewProvider(cfg, &frontend)
	if err != nil {
		return err
	}

	if *generateDb {
		fixtures.InitializeTable(prv, SCINNA_VERSION, *forceGenerateDb)
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
		firstUserInvite, err := prv.GenerateUID()
		if err != nil {
			return err
		}
		code, err := prv.Dal.Registration.GenerateInviteIfNeeded(firstUserInvite)
		if err != nil {
			return err
		}
		if code != "NONE" {
			log.InfoAlwaysShown("Your server is invite only and there are no users registered.")
			log.InfoAlwaysShown("Please use this code to create your admin account: " + code)
		}
	}

	router := mux.NewRouter().StrictSlash(false)

	routes.WebApp(prv, router)

	api := router.PathPrefix("/api").Subrouter()
	routes.Authentication(prv, api.PathPrefix("/auth").Subrouter())
	routes.Accounts(prv, api.PathPrefix("/account").Subrouter())
	routes.Upload(prv, api.PathPrefix("/upload").Subrouter())
	routes.Browser(prv, api.PathPrefix("/browse").Subrouter())

	// Last one (Matching the media_id)
	routes.Medias(prv, router.PathPrefix("").Subrouter())

	srv := &http.Server{
		Handler:           router,
		Addr:              cfg.ListeningAddr,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	return srv.ListenAndServe()
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
