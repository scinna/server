package config

type Config struct {
	ConfigSMTP SMTP
	ConfigDB   DB
	WebPath    string
	WebPort    int
	MediaPath  string
	RegistrationAllowed bool
}

/** SMTP represents the configuration for the database **/
type SMTP struct {
	Enabled bool
	ConnectionType string

	Hostname string
	Port     int
	Username string
	Password string
	Sender   string
}

/** DB represents the configuration for the database **/
type DB struct {
	Dbms     string

	Hostname string
	Port     string
	Username string
	Password string
	Database string
}
