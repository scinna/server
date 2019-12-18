package configuration

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Load loads the configuration from the file
func Load(path string) Configuration {

	if len(path) == 0 {
		path = FindPath()
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println(path + ": This file cannot be found!\nGenerate one withe the command `scinna -genconf > /etc/scinna.yml`")
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var cfg Configuration
	err = yaml.Unmarshal(bts, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

// FindPath finds the best match for the config file
func FindPath() string {
	if _, err := os.Stat("scinna.yml"); os.IsExist(err) {
		return "scinna.yml"
	}

	if _, err := os.Stat("config.yml"); os.IsExist(err) {
		return "config.yml"
	}

	return "/etc/scinna.yml"
}

// Configuration represents the current config for the server
type Configuration struct {
	PostgresDSN         string `yaml:"PostgresDSN"`
	IDAlphabet          string `yaml:"IdAlphabet"`
	IDSize              int    `yaml:"IdSize"`
	WebURL              string `yaml:"WebURL"`
	WebPort             string `yaml:"WebPort"`
	PicturePath         string `yaml:"PicturePath"`
	HeaderIPField       string `yaml:"HeaderIPField"`
	SMTPSender          string `yaml:"SMTPSender"` // @TODO make this a category or something like that
	SMTPHost            string `yaml:"SMTPHost"`
	SMTPUser            string `yaml:"SMTPUser"`
	SMTPPass            string `yaml:"SMTPPass"`
	RegistrationAllowed string `yaml:"RegistrationAllowed"`
	RateLimiting        int    `yaml:"RateLimiting"`
}
