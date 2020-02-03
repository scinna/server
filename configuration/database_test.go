package configuration

import "testing"

func TestDsnSpecialChars(t *testing.T) {
	dbConfig := DBConfig{
		Dbms:     "pgsql",
		Hostname: "localhost",
		Port:     "1234",
		Username: "myUser",
		Password: "my_Password%@/ ",
		Database: "myDb",
	}

	_, dsn := dbConfig.GetDsn()

	expected := "postgres://myUser:my_Password%25%40%2F+@localhost:1234/myDb"

	if dsn != expected {
		t.Errorf("DSN Special chars encoding failed: Wanted \"%v\", received \"%v\"", expected, dsn)
	}
}

func TestDsnDriverPostgres(t *testing.T) {
	dbConfig := DBConfig{
		Dbms:     "pgsql",
		Hostname: "localhost",
		Port:     "1234",
		Username: "myUser",
		Password: "my_Password%@/ ",
		Database: "myDb",
	}

	expected := "postgres"

	driver, _ := dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}

	dbConfig.Dbms = "postgres"
	driver, _ = dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}

	dbConfig.Dbms = "postgresql"
	driver, _ = dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}
}

func TestDsnDriverMysql(t *testing.T) {
	dbConfig := DBConfig{
		Dbms:     "mysql",
		Hostname: "localhost",
		Port:     "1234",
		Username: "myUser",
		Password: "my_Password%@/ ",
		Database: "myDb",
	}

	expected := "mysql"

	driver, _ := dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}

}

func TestDsnDriverSqlite(t *testing.T) {
	dbConfig := DBConfig{
		Dbms:     "sqlite3",
		Hostname: "localhost",
		Port:     "1234",
		Username: "myUser",
		Password: "my_Password%@/ ",
		Database: "myDb",
	}

	expected := "sqlite3"

	driver, _ := dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}

	dbConfig.Dbms = "sqlite"
	driver, _ = dbConfig.GetDsn()
	if driver != expected {
		t.Errorf("DSN driver failed: Wanted \"%v\", received \"%v\"", expected, driver)
	}
}
