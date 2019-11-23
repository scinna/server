package routes

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/services"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	CurrentUser model.AppUser
	Token       string
}

func LoginRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var rc LoginRequest 

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

		valid, err := prv.VerifyPassword(rc.Password, u.Password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if valid {
			token, err := auth.GenerateJWT(prv, u)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			response := &LoginResponse {
				CurrentUser: u,
				Token: string(token),
			}

			sResponse, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write(sResponse)

			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func RefreshRoute (prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("RefreshRoute - To be implemented"))
	}
}