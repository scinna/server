package config

import (
	"encoding/json"
	"errors"
	"github.com/scinna/server/log"
	"io/ioutil"
	"os"
)

type Config struct {
	Mail            SMTP
	Database        DB
	Registration    Registration
	WebURL          string
	ListeningAddr   string
	MediaPath       string
	CustomLogoWide  string
	CustomLogoSmall string
	CustomBranding  string
}

func (c *Config) Validate() error {
	err := []error{
		c.Mail.Validate(),
		c.Database.Validate(),
		c.Registration.Validate(),
	}

	if len(c.WebURL) == 0 {
		err = append(err, errors.New("WebURL can't be empty"))
	}

	if len(c.ListeningAddr) == 0 {
		err = append(err, errors.New("ListeningAddr can't be empty"))
	}

	if len(c.MediaPath) == 0 {
		err = append(err, errors.New("MediaPath can't be empty"))
	}

	return combineErrors(err...)
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
	if err != nil {
		return nil, err
	}

	err = cfg.Validate()

	return &cfg, err
}
