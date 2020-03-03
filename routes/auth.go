package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/scinna/server/auth"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
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
			serrors.ErrorBadRequest.Write(w)
			return
		}

		var rc loginRequest

		err = json.Unmarshal(body, &rc)
		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		u, err := dal.GetUser(prv, rc.Username)
		if err != nil {
			if err == serrors.ErrorUserNotFound {
				serrors.ErrorInvalidCredentials.Write(w)
				return
			}
			serrors.WriteError(w, err)
			return
		}

		err = dal.ReachedMaxAttempts(prv, u)
		if serrors.WriteError(w, err) {
			return
		}

		if !u.Validated {
			serrors.ErrorNotValidated.Write(w)
			return
		}

		valid, err := prv.VerifyPassword(rc.Password, u.Password)
		if err != nil {
			serrors.WriteLoggableError(w, err)
			return
		}

		if valid {
			token, err := auth.GenerateToken(prv, utils.ReadUserIP(prv.Config.HeaderIPField, r), u)
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
		serrors.ErrorInvalidCredentials.Write(w)
		dal.InsertFailedLoginAttempt(prv, u, utils.ReadUserIP(prv.Config.HeaderIPField, r))
	}
}

// CheckTokenRoute is the route that lets the user authenticate: /auth/login
func CheckTokenRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		userMarshalled, err := json.Marshal(user)
		if err != nil {
			serrors.ErrorTokenNotFound.Write(w)
			return
		}

		w.Write(userMarshalled)
	}
}

// GetTokensRoute sends all the user's token
func GetTokensRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		tokens, err := dal.ListTokens(prv, user)
		if serrors.WriteError(w, err) {
			return
		}

		tks, err := json.Marshal(tokens)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(tks)

	}
}

// RevokeTokenRoute revoke a token
func RevokeTokenRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id := params["TOKEN_ID"]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			serrors.ErrorBadRequest.Write(w)
			return
		}

		u, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		err = dal.RevokeToken(prv, idInt, *u.ID)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusGone)
	}
}
