package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/scinna/server/auth"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

// IsRegisterAvailableRoute is 200 when you can register, and 403 when you cant
func IsRegisterAvailableRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonRsp := fmt.Sprintf("{ \"Registration\": \"%v\"}", prv.Config.RegistrationAllowed)
		w.Write([]byte(jsonRsp))
	}
}

// RegisterRequest reprensent the request that let users register
type RegisterRequest struct {
	Username   string
	Email      string
	Password   string
	InviteCode string
}

// RegisterRoute lets someone register on the server
func RegisterRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ToLower(prv.Config.RegistrationAllowed) == "false" {
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

		var invitedBy int64 = -1
		if strings.ToLower(prv.Config.RegistrationAllowed) == "invite" {
			if len(rc.InviteCode) == 0 {
				serrors.ErrorInviteOnly.Write(w)
				return
			}

			code, err := dal.GetInviteCode(prv, rc.InviteCode)
			if serrors.WriteError(w, err) {
				return
			}
			invitedBy = *code.GeneratedBy.ID
		}

		if len(rc.Username) == 0 || rc.Username == "me" {
			serrors.ErrorRegBadUsername.Write(w)
			return
		}

		if !prv.Config.Mail.IsEmail.MatchString(rc.Email) {
			serrors.ErrorRegBadEmail.Write(w)
			return
		}

		_, err = dal.RegisterUser(prv, rc.Username, rc.Password, rc.Email, invitedBy)
		if serrors.WriteError(w, err) {
			return
		}

		w.WriteHeader(http.StatusCreated)

		serrors.WriteLoggableError(w, dal.RevokeInviteCode(prv, rc.InviteCode))
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

		prv.Templates.ExecuteTemplate(w, "validation_mail.tmpl", vrr)

	}
}

// GenerateInviteRoute lets administrators generate invitations for the server
func GenerateInviteRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, err := auth.ValidateRequest(prv, w, r)
		if serrors.WriteError(w, err) {
			return
		}

		if user.Role != model.UserRoleAdmin {
			serrors.WriteError(w, serrors.ErrorNotAdmin)
			return
		}

		code, err := dal.GenerateInviteCode(prv, &user)

		if serrors.WriteLoggableError(w, err) {
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("{ \"InvitationCode\": \"" + code + "\" }"))
	}
}
