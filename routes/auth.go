package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/utils"
)

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	CurrentUser model.AppUser
	Token       string
}

// LoginRoute is the route that lets the user authenticate: /auth/login
func LoginRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var rc loginRequest

		err = json.Unmarshal(body, &rc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := dal.GetUser(prv, rc.Username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		hasReachedMaxAttempts, err := dal.ReachedMaxAttempts(prv, u)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if hasReachedMaxAttempts {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		if !u.Validated {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		valid, err := prv.VerifyPassword(rc.Password, u.Password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if valid {
			token, err := auth.GenerateToken(prv, utils.ReadUserIP(prv.HeaderIPField, r), u)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			response := &loginResponse{
				CurrentUser: u,
				Token:       string(token),
			}

			sResponse, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write(sResponse)

			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		dal.InsertFailedLoginAttempt(prv, u, utils.ReadUserIP(prv.HeaderIPField, r))
	}
}

// IsRegisterAvailableRoute is 200 when you can register, and 403 when you cant
func IsRegisterAvailableRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !prv.RegistrationAllowed {
			w.WriteHeader(http.StatusForbidden)
		}

		w.Write([]byte(`
		{
			"RegisterAllowed": ` + strconv.FormatBool(prv.RegistrationAllowed) + `
		}
		`))
	}
}

// RegisterRequest reprensent the request that let users register
type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

// RegisterRoute lets someone register on the server
func RegisterRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if prv.RegistrationAllowed {
			serrors.ErrorRegDisabled.Write(w)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		var rc RegisterRequest

		err = json.Unmarshal(body, &rc)
		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		if len(rc.Username) == 0 || rc.Username == "me" {
			serrors.ErrorRegBadUsername.Write(w)
			return
		}

		if !prv.Mail.IsEmail.MatchString(rc.Email) {
			serrors.ErrorRegBadEmail.Write(w)
			return
		}

		_, err = dal.RegisterUser(prv, rc.Username, rc.Password, rc.Email)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

type validateRouteResp struct {
	Title            string
	Validated        bool
	AlreadyValidated bool
	NotFound         bool
	ErrMsg           string
}

// ValidateUserRoute lets someone register on the server
func ValidateUserRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		token := params["VALIDATION_TOKEN"]

		err := dal.ValidateUser(prv, token)

		vrr := validateRouteResp{
			Title:            "Activating user - Scinna",
			Validated:        err == nil,
			AlreadyValidated: err == serrors.ErrorAlreadyValidated,
			NotFound:         err == serrors.ErrorNoAccountValidation,
		}

		if err != nil {
			vrr.ErrMsg = err.Error()
		}

		prv.Templates.ExecuteTemplate(w, "layout", vrr)

	}
}
