package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/scinna/server/configuration"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/services"
	"github.com/scinna/server/wizard"

	_ "github.com/lib/pq"
)

func main() {

	var configFlag = flag.Bool("genconf", false, "Show a config example")

	// If not specified, should also check in the env vars
	var configPath = flag.String("config", "", "Specify a path to load the config from")
	var port = flag.Int("port", -1, "Specify on which port to listen")

	flag.Parse()

	/** We are going to keep this even though we have a web-installer for those who might want a headless setup **/
	if *configFlag {
		configuration.PrintExample()
		return
	}

	fmt.Println("Scinna Server - V1")
	if configuration.HasConfig(configPath) {
		StartFullServer(configPath, port)
	} else {
		if *port > 0 {
			fmt.Println("This looks like the first startup.")
			fmt.Println("Launching the setup on port " + strconv.Itoa(*port))

			wizard.RunSetup(port)
		} else {
			fmt.Println("No port set! Please specify it with -port [port]")
			os.Exit(1)
		}
	}
}

// StartFullServer runs the full scinna server
func StartFullServer(configPath *string, port *int) {
	cfg, exists, err := configuration.Load(*configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !exists {
		fmt.Println("No config file found! Shouldn't be possible.")
		os.Exit(1)
	}

	if *port > 0 {
		cfg.WebPort = strconv.Itoa(*port)
	}

	prv, err := services.New(cfg)
	prv.Init()
	defer prv.Db.Close()

	if err != nil {
		panic(err)
	}

	// Every 15 minutes, we clean up users older than 24h who have not validated their accounts
	go func(prv *services.Provider) {
		for {
			dal.CleanupUsers(prv)
			time.Sleep(15 * time.Minute)
		}
	}(prv)

	RunServer(prv)
}
