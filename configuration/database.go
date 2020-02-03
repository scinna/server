package configuration

import (
	"fmt"
	"net/url"
	"strings"
)

// DBConfig represents the login infos about a database
type DBConfig struct {
	Dbms string `json:"dbms"`

	Hostname string `json:"db_host"`
	Port     string `json:"db_port"`

	Username string `json:"db_username"`
	Password string `json:"db_password"`
	Database string `json:"db_database"`

	Filepath string `json:"db_path"`
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
