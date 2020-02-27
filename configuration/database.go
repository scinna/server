package configuration

import (
	"fmt"
	"net/url"
	"strings"
)

// DBConfig represents the login infos about a database
type DBConfig struct {
	Dbms string `json:"Dbms"`

	Hostname string `json:"Hostname"`
	Port     string `json:"Port"`

	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`

	Filepath string `json:"Path"`
}

// GetDsn translate a configuration to a DSN ready to be used
func (cfg *DBConfig) GetDsn() (driver string, dsn string) {
	dbms := strings.ToLower(cfg.Dbms)

	if dbms == "sqlite3" || dbms == "sqlite" {
		driver = "sqlite3"
		dsn = cfg.Filepath
	} else if dbms == "mysql" {
		driver = "mysql"
		dsn = "mysql://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Hostname + ":" + cfg.Port + "/" + cfg.Database
	} else if dbms == "pgsql" || dbms == "postgres" || dbms == "postgresql" {
		driver = "postgres"
		dsn = "postgres://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Hostname + ":" + cfg.Port + "/" + cfg.Database
	} else {
		fmt.Println("- No matching database found: " + cfg.Dbms)
	}

	return driver, dsn
}
