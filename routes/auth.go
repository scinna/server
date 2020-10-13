package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/log"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/requests"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/utils"
	"html/template"
	"net/http"
	"strings"
)

func Authentication(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/register", findRegistrationType(prv)).Methods(http.MethodGet)
	r.HandleFunc("/register", register(prv)).Methods(http.MethodPost)
	r.HandleFunc("/register/{validation_code}", validateAccount(prv)).Methods(http.MethodGet)

	r.HandleFunc("/login", authenticate(prv)).Methods(http.MethodPost)
}

func findRegistrationType(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`{ "Registration": %v, "Validation": "%v" }`, prv.Config.Registration.Allowed, prv.Config.Registration.Validation)))
	}
}

func register(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var registerBody requests.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&registerBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// If the server requires an invite code and is not given, we kick the user
		if !prv.Config.Registration.Allowed && len(registerBody.InviteCode) == 0 {
			serrors.ErrorBadInviteCode.Write(w)
			return
		}

		if len(registerBody.Username) == 0 || len(registerBody.Email) == 0 || len(registerBody.Password) == 0 || strings.ToLower(registerBody.Username) == "me" {
			serrors.ErrorInvalidRegistration.Write(w)
			return
		}

		// Check if the invite code exists
		invite := dal.FindInvite(prv, registerBody.InviteCode)
		if !prv.Config.Registration.Allowed && (invite == nil || invite.Used) {
			if invite == nil {
				serrors.ErrorBadInviteCode.Write(w)
			} else if invite.Used {
				serrors.ErrorInviteUsed.Write(w)
			}
			return
		}

		// Register the user
		valcode, err := dal.RegisterUser(prv, &registerBody)
		if err != nil {
			if err, ok := err.(*pq.Error); ok {
				if err.Constraint == "scinna_user_user_name_key"  {
					serrors.ErrorUserExists.Write(w)
				} else if err.Constraint == "scinna_user_user_email_key" {
					serrors.ErrorEmailExists.Write(w)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}

				return
			}
		}

		// Send the mail
		if prv.Config.Registration.Validation == "email" && prv.Config.ConfigSMTP.Enabled {
			_, err := prv.SendValidationMail(registerBody.Email, valcode)
			if err != nil {
				log.Warn(err.Error())
			}
		}

		// Set the invite code to used
		if invite != nil {
			dal.DisableInvite(prv, invite)
		}

		// Send the response
		if prv.Config.Registration.Validation == "admin" {
			serrors.UserNeedValidationAdmin.Write(w)
			return
		} else if prv.Config.Registration.Validation == "email" {
			serrors.UserNeedValidationEmail.Write(w)
			return
		} else {
			serrors.UserRegistered.Write(w)
		}
	}
}

func validateAccount(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")

		validationCode := mux.Vars(r)["validation_code"]
		user := dal.ValidateUser(prv, validationCode)

		t := template.Must(template.ParseFiles("templates/validated.html"))
		t.Execute(w, struct {
			Username string
		} {
			Username: user,
		})
	}
}

func authenticate(prv *services.Provider) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRq requests.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&authRq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := dal.GetUserFromUsername(prv, authRq.Username)
		if err != nil {
			serrors.InvalidUsernameOrPassword.Write(w)
			return
		}

		if !user.Validated {
			serrors.AccountNotValidated.Write(w)
			return
		}

		isValid, err := prv.VerifyPassword(authRq.Password, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if isValid {
			token, err := dal.Login(prv, user, utils.IPForRequest(prv.Config, r))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			authResp := dto.AuthDto{
				User:  user,
				Token: token,
			}

			resp, _ := json.Marshal(authResp)
			w.Write(resp)

			return
		}

		serrors.InvalidUsernameOrPassword.Write(w)
	}
}