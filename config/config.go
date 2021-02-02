package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/scinna/server/log"
)

type Config struct {
	Mail          SMTP
	Database      DB
	Registration  Registration
	WebURL        string
	ListeningAddr string
	MediaPath     string
}

/** SMTP represents the configuration for the database **/
type SMTP struct {
	Enabled        bool
	ConnectionType string

	Hostname string
	Port     int
	Username string
	Password string
	Sender   string
}

/** DB represents the configuration for the database **/
type DB struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

type Registration struct {
	HasUser    bool
	Allowed    bool
	Validation string
}

func Load() (*Config, error) {
	if _, err := os.Stat("./config.json"); !os.IsNotExist(err) {
		log.Info("Using config file in the current folder")
		return loadFile("./config.json")
	}

	if _, err := os.Stat("/etc/scinna/config.json"); !os.IsNotExist(err) {
		log.Info("Using config file in /etc/scinna/config.json")
		return loadFile("/etc/scinna/config.json")
	}

	return nil, errors.New("can't find the config file")
}

func loadFile(file string) (*Config, error) {
	var cfg Config

	jsonFile, err := os.Open(file)
	if err != nil {
		return &cfg, err
	}
	defer jsonFile.Close()

	bytesInFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &cfg, err
	}

	err = json.Unmarshal(bytesInFile, &cfg)

	cfg.Registration.Validation = strings.ToLower(cfg.Registration.Validation)
	if err == nil && cfg.Registration.Validation != "open" && cfg.Registration.Validation != "email" && cfg.Registration.Validation != "admin" {
		err = errors.New("Registration validation must match one of these values: \"open\", \"email\" or \"admin\"")
	}

	return &cfg, err
}
