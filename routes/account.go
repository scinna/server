package routes

import (
	"encoding/json"
	"fmt"
	"github.com/scinna/server/forms"
	"github.com/scinna/server/requests"
	"github.com/scinna/server/translations"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
)

func Accounts(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.LoggedInMiddleware(prv))
	r.Use(middlewares.Json)

	r.HandleFunc("", fetchAccountInfos(prv)).Methods(http.MethodGet)
	r.HandleFunc("", updateAccountInfos(prv)).Methods(http.MethodPut)
	r.HandleFunc("/tokens", fetchTokens(prv)).Methods(http.MethodGet)
	r.HandleFunc("/tokens/{token}", revokeToken(prv)).Methods(http.MethodDelete)
}

func fetchAccountInfos(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		u, _ := json.Marshal(user)
		_, _ = w.Write(u)
	}
}

func fetchTokens(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		tokens, err := prv.Dal.User.FetchUserTokens(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		u, _ := json.Marshal(tokens)
		_, _ = w.Write(u)
	}
}

func revokeToken(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)
		token := mux.Vars(r)["token"]

		if !prv.Dal.User.TokenBelongsToUser(user, token) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		revokedAt, err := prv.Dal.User.RevokeToken(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusGone)
		bytes, _ := json.Marshal(map[string]interface{}{
			"RevokedAt": revokedAt,
		})
		_, _ = w.Write(bytes)
	}
}

func updateAccountInfos(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(*models.User)

		rq := requests.EditAccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&rq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(rq.CurrentPassword) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isValid, err := prv.VerifyPassword(rq.CurrentPassword, user.Password)
		if err != nil || !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			bytes, _ := json.Marshal(struct {
				Violations map[string]string
			}{
				Violations: map[string]string{
					"CurrentPassword": translations.T(r, "errors.login.invalid_password"),
				},
			})
			w.Write(bytes)
			return
		}

		if len(rq.Password) > 0 {
			hashedPassword, err := prv.HashPassword(rq.Password)
			if err != nil {
				// @TODO before 1.0 => Make "loggable errors" that writes to the stdout/stderr
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Could not hash password")
				return
			}

			rq.Password = hashedPassword
			prv.Dal.User.UpdatePassword(user, rq.Password)
		}

		if rq.Email != user.Email {
			// @TODO: Validate that the email contains an @

			violations := []forms.Constraint{
				forms.ConstraintUniqueString(prv, "Email", "SCINNA_USER", "user_email", rq.Email, translations.T(r, "errors.registration.email_exists")),
			}

			if forms.HasViolations(violations, w){
				return
			}

			if false && prv.Config.Mail.Enabled {
				// @TODO Use a SQL trigger to detect when we set the validated flag to false and automagically
				// generate a new validation code
				// @TODO Send confirm mail
				prv.Dal.User.UpdateEmail(user, rq.Email, false)

			} else {
				prv.Dal.User.UpdateEmail(user, rq.Email, true)
			}

			// We don't return anything as when we'll switch to JWT
			// based auth we'll simply invalidate the JWT token client-side and he'll refresh
			// using new info.
		}
	}
}