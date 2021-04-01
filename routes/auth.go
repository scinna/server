package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/goware/emailx"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/forms"
	"github.com/scinna/server/log"
	"github.com/scinna/server/middlewares"
	"github.com/scinna/server/models"
	"github.com/scinna/server/requests"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"github.com/scinna/server/translations"
	"github.com/scinna/server/utils"
	"net/http"
)

func Authentication(prv *services.Provider, r *mux.Router) {
	r.Use(middlewares.Json)

	/**
	 * Maybe we will switch to JWT at some point, but not now, there are few enough users that it doesn't matter
	 */
	r.Handle("", prv.RateLimiting.LoginMiddleware.Handle(authenticate(prv))).Methods(http.MethodPost)
	r.Handle("", middlewares.LoggedInMiddleware(prv)(logout(prv))).Methods(http.MethodDelete)

	r.Handle("/register", prv.RateLimiting.RegistrationMiddleware.Handle(register(prv))).Methods(http.MethodPost)
	r.HandleFunc("/register/{validation_code}", validateAccount(prv)).Methods(http.MethodGet)
	r.Handle("/forgotten_password", prv.RateLimiting.ForgottenPasswordMiddleware.Handle(forgottenPassword(prv))).Methods(http.MethodGet)
	r.HandleFunc("/forgotten_password/{validation_code}", setNewPassword(prv)).Methods(http.MethodPost)
}

func register(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registerBody requests.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&registerBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// If the server requires an invite code and is not given, we kick the user
		if !prv.Config.Registration.Allowed && len(registerBody.InviteCode) == 0 {
			serrors.ErrorBadInviteCode.Write(w, r)
			return
		}

		if len(registerBody.Username) == 0 || len(registerBody.Email) == 0 || len(registerBody.Password) == 0 {
			serrors.ErrorInvalidRegistration.Write(w, r)
			return
		}

		// Check if the invite code exists
		var invite *models.InviteCode
		if !prv.Config.Registration.Allowed {
			invite = prv.Dal.Registration.FindInvite(registerBody.InviteCode)
			if invite == nil {
				serrors.ErrorBadInviteCode.Write(w, r)
				return
			} else if invite.Used {
				serrors.ErrorInviteUsed.Write(w, r)
				return
			}
		}

		// meh this library, might change later or just check that there is an @ and no spaces in the str
		err = emailx.Validate(registerBody.Email)
		if err != nil {
			serrors.InvalidEmail.Write(w, r)
			return
		}

		// Maybe one day I'll have the motivation to rewrite this with golang's tag
		violations := []forms.Constraint{
			forms.ConstraintUniqueString(prv, "Email", "SCINNA_USER", "user_email", registerBody.Email, translations.T(r, "errors.registration.email_exists")),
			forms.ConstraintUniqueString(prv, "Username", "SCINNA_USER", "user_name", registerBody.Username, translations.T(r, "errors.registration.username_exists")),
		}

		if forms.HasViolations(violations, w) {
			return
		}

		hashedPassword, err := prv.HashPassword(registerBody.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		registerBody.HashedPassword = hashedPassword

		// Register the user
		valCode, err := prv.Dal.Registration.RegisterUser(&registerBody, prv.Config.Registration.Validation == "open")
		if err != nil {
			// @TODO do something better
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the mail
		if prv.Config.Mail.Enabled && prv.Config.Registration.Validation == "email" {
			_, err := prv.SendValidationMail(r, registerBody.Email, valCode)
			if err != nil {
				// @TODO Add a warning for the user
				log.Warn(err.Error())
			}
		}

		// Set the invite code to used
		if invite != nil {
			prv.Dal.Registration.DisableInvite(invite)
		}

		// Send the response
		switch prv.Config.Registration.Validation {
		case "admin":
			serrors.UserNeedValidationAdmin.Write(w, r)
		case "email":
			serrors.UserNeedValidationEmail.Write(w, r)
		default:
			serrors.UserRegistered.Write(w, r)
		}
	}
}

func validateAccount(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validationCode := mux.Vars(r)["validation_code"]
		user := prv.Dal.Registration.ValidateUser(validationCode)

		if len(user) == 0 {
			serrors.InvalidValidationCode.Write(w, r)
			return
		}

		_, _ = w.Write([]byte(fmt.Sprintf(`{ "username": "%s" }`, user)))
	}
}

func authenticate(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRq requests.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&authRq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := prv.Dal.User.GetUserFromUsername(authRq.Username)
		if err != nil {
			serrors.InvalidUsernameOrPassword.Write(w, r)
			return
		}

		if !user.Validated {
			serrors.AccountNotValidated.Write(w, r)
			return
		}

		isValid, err := prv.VerifyPassword(authRq.Password, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !isValid {
			serrors.InvalidUsernameOrPassword.Write(w, r)
			return
		}

		token, err := prv.Dal.User.Login(user, utils.IPForRequest(r))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authResp := dto.AuthDto{
			User:  user,
			Token: token,
		}

		resp, _ := json.Marshal(authResp)
		_, _ = w.Write(resp)
	}
}

func logout(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("token").(string)
		prv.Dal.User.RevokeToken(token)
	}
}

func forgottenPassword(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !prv.Config.Mail.Enabled {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		sentIfExists, err := json.Marshal(struct{
			Message string
		}{
			Message: translations.T(r, "forgotten_password.sent_if_exists"),
		})
		if serrors.WriteError(w, r, err) {
			return
		}

		username := r.FormValue("username")
		user, err := prv.Dal.User.GetUserFromUsername(username)
		if err != nil {
			if err == sql.ErrNoRows {
				_, _ = w.Write(sentIfExists)
				return
			}

			serrors.WriteError(w, r, err)
			return
		}

		valCode, err := prv.Dal.User.RequestPasswordReset(user)
		if err != nil {
			// @TODO do something better
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = prv.SendForgottenPasswordMail(r, user.Email, valCode)
		if err != nil {
			// @TODO Add a warning for the user
			log.Warn(err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(sentIfExists)
	}
}

func setNewPassword(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
