package fixtures

import (
	"fmt"
	"github.com/scinna/server/log"

	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
)

func InitializeTable(prv *services.Provider, scinnaVersion string, force bool) {
	// Install required plugins
	plugins := []string{
		"citext",
		"pgcrypto",
	}

	log.InfoAlwaysShown("Creating the database...")

	if !force {
		log.InfoAlwaysShown("Executing the creation script will remove ALL TABLES AND ALL DATA.")
		log.InfoAlwaysShown("Are you sure you want to proceed ? (y / N)")

		// @TODO read user input
		// if != Y & != y => os.exit(1)
	}

	for _, p := range plugins {
		_, err := prv.DB.Exec(fmt.Sprintf("CREATE EXTENSION IF NOT EXISTS %v;", p))
		if err != nil {
			panic(err)
		}
	}

	createVersionTable(prv, scinnaVersion)

	tables := []models.DatabaseModel{
		models.User{},
		models.Collection{},
		models.AuthToken{},
		models.InviteCode{},
		models.Media{},
	}

	// Dropping old tables
	for i := len(tables) - 1; i >= 0; i-- {
		_, e := prv.DB.Exec(`DROP TABLE IF EXISTS ` + tables[i].GetTableName())
		if e != nil {
			panic(e)
		}
	}

	// Re-creating tables
	for _, t := range tables {
		// No need to continue to run if there is an error, since this will never be called on a prod machine
		_, e := prv.DB.Exec(t.GenerateTable())
		if e != nil {
			panic(e)
		}
		log.InfoAlwaysShown(fmt.Sprintf("\t- Table %v created.", t.GetTableName()))
	}

	// Inserting default user
	pwd, err := prv.HashPassword("admin")
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Name:           "admin",
		Email:          "admin@scinna.app",
		Password:       pwd,
		Validated:      true,
	}

	log.InfoAlwaysShown("\t- Inserting default user (admin:admin)")
	err = prv.Dal.User.InsertUser(user)
	if err != nil {
		panic(err)
	}
}

func createVersionTable(prv *services.Provider, version string) {
	_, err := prv.DB.Exec(`
	DROP TABLE IF EXISTS DBVERSION;
	CREATE TABLE DBVERSION
	(
		VERSION VARCHAR
	);`)

	if err != nil {
		panic(err)
	}

	_, err = prv.DB.Exec(`INSERT INTO DBVERSION (VERSION) VALUES ($1);`, version)
	if err != nil {
		panic(err)
	}
}
