package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/oxodao/scinna/configuration"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/services"

	_ "github.com/lib/pq"
)

func main() {

	var configFlag = flag.Bool("genconf", false, "Show a config example")
	var configPath = flag.String("config", "", "Specify a path to load the config from")

	flag.Parse()

	if *configFlag {
		configuration.PrintExample()
		return
	}

	fmt.Println("Scinna Server - V1")

	cfg := configuration.Load(*configPath)

	prv := services.New(cfg)
	defer prv.Db.Close()

	// Every 15 minutes, we clean up users older than 24h who have not validated their accounts
	go func(prv *services.Provider) {
		for {
			dal.CleanupUsers(prv)
			time.Sleep(15 * time.Minute)
		}
	}(prv)

	RunServer(prv)
}
