package routes

import (
	"encoding/json"
	"errors"
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

var currentConfig configuration.Configuration = configuration.Configuration{}

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
	if err != nil || db == nil || db.Ping() != nil {
		w.Write([]byte("{ \"IsValid\": false }"))
		return
	}
	db.Close()

	currentConfig.Database = params

	w.Write([]byte("{ \"IsValid\": true }"))
}

// TestSMTPConfigRoute is the SPA that let the user set the server up
func TestSMTPConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params configuration.MailConfig = configuration.MailConfig{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !params.Enabled {
		currentConfig.Mail = params
		w.Write([]byte("{ \"IsValid\": true }"))
		return
	}

	currentConfig.Mail = params

	wasSent, err := params.SendMail(params.TestReceiver, "Scinna test email", "This email was sent by Scinna. It's a test to check your SMTP settings.")

	if !wasSent || err != nil {
		if err == nil {
			err = errors.New("The message was not sent")
		}

		msg := struct {
			IsValid bool
			Message error
		}{
			IsValid: false,
			Message: err,
		}

		msgJSON, err := json.Marshal(msg)
		if err != nil {
			msgJSON = []byte("{\"IsValid\": false, \"Message\": \"JSON encoder went wrong.\"}")
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(msgJSON)
		return
	}

	w.Write([]byte("{ \"IsValid\": true }"))
}

// CreateAdminRoute is called when the user first setup the server to create the original account
func CreateAdminRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Coucou")
}
