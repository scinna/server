package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

// Configuration represents the current config for the server
type Configuration struct {
	Database            DBConfig   `yaml:"Database"`
	Mail                MailConfig `yaml:"Mail"`
	IDAlphabet          string     `yaml:"IdAlphabet"`
	IDSize              int        `yaml:"IdSize"`
	WebURL              string     `yaml:"WebURL"`
	WebPort             string     `yaml:"WebPort"`
	PicturePath         string     `yaml:"PicturePath"`
	HeaderIPField       string     `yaml:"HeaderIPField"`
	RegistrationAllowed string     `yaml:"RegistrationAllowed"`
	RateLimiting        int        `yaml:"RateLimiting"`
}

// Meh global variable >:(
var lastTriedPath string = ""

// HasConfig checks if the config file exists
func HasConfig(path *string) bool {
	if len(*path) == 0 {
		osPath := FindPath()
		path = &osPath
	}

	lastTriedPath = *path

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Load loads the configuration from the file, returns the config, whether the file existed and an error
func Load(path string) (*Configuration, bool, error) {
	if len(path) == 0 {
		path = FindPath()
	}

	if !HasConfig(&path) {
		return nil, false, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, true, err
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, true, err
	}

	var cfg Configuration
	err = yaml.Unmarshal(bts, &cfg)
	if err != nil {
		return nil, true, err
	}

	reg, err := regexp.Compile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	if err != nil {
		// @TODO Log it really (Redmine #109)
		fmt.Println(err)
	}

	cfg.Mail.IsEmail = reg

	return &cfg, true, err
}

// SaveConfig saves the configuration to the file
func SaveConfig(cfg *Configuration) error {
	cfgBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(lastTriedPath, cfgBytes, 0644)
}

// FindPath finds the best match for the config file
func FindPath() string {
	if _, err := os.Stat("scinna.yml"); os.IsExist(err) {
		return "scinna.yml"
	}

	home, err := os.UserHomeDir()
	if err == nil {
		return home + "/scinna.yml"
	}

	// Default path, meh
	//@TODO: Add the findpath equivalent for Windows (%APPDATA%) and OSX (/usr/local/etc/)
	return "/etc/scinna.yml"
}
