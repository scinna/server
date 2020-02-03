package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/configuration"
	"github.com/scinna/server/services"
)

// IndexRoute is the index endpoint, the one displaying the react webapp
func IndexRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This will be the react app for Scinna picture server"))
	}
}

/**

	These are routes for the first-launch setup

**/

var currentConfig configuration.Configuration

// SaveConfigRoute is the SPA that let the user set the server up
func SaveConfigRoute(w http.ResponseWriter, r *http.Request) {

}

// TestDatabaseConfigRoute is the SPA that let the user set the server up
func TestDatabaseConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params configuration.DBConfig

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	driver, dsn := params.GetDsn()
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		w.Write([]byte("{ \"IsValid\": false }"))
		return
	}
	db.Close()

	w.Write([]byte("{ \"IsValid\": true }"))
}

type smtpTestParams struct {
	MailDisabled bool
	Hostname     string `json:"smtp_host"`
	Port         string `json:"smtp_port"`
	Username     string `json:"smtp_username"`
	Password     string `json:"smtp_password"`
}

// TestSMTPConfigRoute is the SPA that let the user set the server up
func TestSMTPConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params smtpTestParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.MailDisabled {
		w.Write([]byte("{ \"IsValid\": true }"))
		return
	}

	// @TODO: Check plus en détail pourquoi ça marche pas
	// ex: timeout, etc... à afficher sur le client
	// Si la connexion fonctionne: stocker dans la constante globale le setting afin qu'il la garde en dernier truc
	if true {
		w.Write([]byte("{ \"IsValid\": true }"))
	} else {
		w.Write([]byte("{ \"IsValid\": false }"))
	}
}

func CreateAdminRoute(w http.ResponseWriter, r *http.Request) {

}
