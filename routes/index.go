package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/scinna/server/configuration"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
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
	They are ONLY available FOR the first-launch setup. Disabled afterwards

**/

var currentConfig configuration.Configuration = configuration.Configuration{}

type errorConfig struct {
	IsValid bool
	Message string
}

func writeError(w http.ResponseWriter, writeHeader bool, err error) {
	if err != nil {
		if writeHeader {
			w.WriteHeader(http.StatusInternalServerError)
		}

		resp, err := json.Marshal(errorConfig{
			IsValid: false,
			Message: err.Error(), // For now, later it will require to compare and produce explicit and clean error message
		})

		if err != nil {
			w.Write([]byte("{ \"IsValid\": false, \"Message\": \"Could not encode JSON error!\" }"))
			return
		}

		w.Write(resp)
	}
}

// TestDatabaseConfigRoute is the route that let the user configure and test the database settings
func TestDatabaseConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params configuration.DBConfig

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, false, err)
		return
	}

	driver, dsn := params.GetDsn()
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		writeError(w, true, err)
		return
	}

	if db == nil {
		writeError(w, true, errors.New("Database is nil")) // @TODO: Find when it can happen and write a corresponding error message
		return
	}

	if db.Ping() != nil {
		writeError(w, true, errors.New("Can't ping the database")) // @TODO: Find when it can happen and write a corresponding error message
		return
	}

	db.Close()

	currentConfig.Database = params

	w.Write([]byte("{ \"IsValid\": true, \"Message\": \"\" }"))
}

// TestSMTPConfigRoute is the route that let the user configure and test the email settings
func TestSMTPConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params configuration.MailConfig = configuration.MailConfig{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, false, err)
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

		writeError(w, true, err)
		return
	}

	w.Write([]byte("{ \"IsValid\": true }"))
}

// ScinnaConfigRoute is the route that let the user configure the scinna settings
func ScinnaConfigRoute(w http.ResponseWriter, r *http.Request) {
	var params configuration.Configuration

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, false, err)
		return
	}

	if len(params.RegistrationAllowed) > 0 && len(params.PicturePath) > 0 && len(params.WebURL) > 0 {
		w.WriteHeader(http.StatusOK)

		params.Database = currentConfig.Database
		params.Mail = currentConfig.Mail

		currentConfig = params

		return
	}

	writeError(w, true, errors.New("Missing fields"))
}

type createAdminRequest struct {
	Username string
	Email    string
	Password string
}

// CreateAdminRoute is the route that create the initial admin account
func CreateAdminRoute(w http.ResponseWriter, r *http.Request) {
	prv, err := services.New(&currentConfig)
	if err != nil {
		writeError(w, true, err)
		return
	}

	prv.Init()

	/**
		@TODO: Initialize database when migration to the ORM is done
	**/

	var params createAdminRequest = createAdminRequest{}

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		writeError(w, true, err)
		return
	}

	hPass, err := prv.HashPassword(params.Password)
	if err != nil {
		writeError(w, true, err)
		return
	}

	rq := ` INSERT INTO APPUSER(USERNAME, EMAIL, PASSWORD, INVITED_BY, VALIDATED, ROLE) 
			VALUES ($1, LOWER($2), $3, NULL, true, $4)`

	_, err = prv.Db.Query(rq, params.Username, params.Email, hPass, model.UserRoleAdmin)
	if err != nil {
		if errPost, ok := err.(*pq.Error); ok {
			/** Should not happen since the setup only shows up when there are no users **/
			if errPost.Code.Name() == serrors.PostgresError["AlreadyExisting"] {
				serrors.WriteError(w, serrors.ErrorRegExistingUser)
				return
			}
		}

		writeError(w, true, serrors.ErrorBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	// Should write the config to the file and restart the server
	// This require the router to not push into the history
	configuration.SaveConfig(&currentConfig)

	go func() {
		// Ugly but IDK how to do it better
		// Should wait until this route has responded then shut off the server
		// Ideally it should also restart itself but this seems a bit complicated
		time.Sleep(5 * time.Second)
		prv.Shutdown()
		os.Exit(0)
	}()
}
