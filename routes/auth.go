package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
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

		valid, err := prv.VerifyPassword(rc.Password, u.Password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if valid {
			token, err := auth.GenerateToken(prv, utils.ReadUserIP(prv, r), u)
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
		dal.InsertFailedLoginAttempt(prv, u, utils.ReadUserIP(prv, r))
	}
}

// RefreshRoute is the route that let the user refresh his JWT token: /auth/refresh
func RefreshRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("RefreshRoute - To be implemented"))
	}
}
