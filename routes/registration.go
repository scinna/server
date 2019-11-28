package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oxodao/scinna/auth"
	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// IsRegisterAvailableRoute is 200 when you can register, and 403 when you cant
func IsRegisterAvailableRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if prv.Config.RegistrationAllowed == 2 {
			w.WriteHeader(http.StatusForbidden)
		}

		jsonRsp := "{ \"Registration\": \""

		switch prv.Config.RegistrationAllowed {
		case 0:
			jsonRsp = jsonRsp + "PUBLIC"
			break

		case 1:
			jsonRsp = jsonRsp + "INVITE"
			break

		case 2:
			jsonRsp = jsonRsp + "PRIVATE"
			break
		}

		jsonRsp = jsonRsp + "\" }"

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
		if prv.Config.RegistrationAllowed == 2 {
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
		if prv.Config.RegistrationAllowed == 1 {
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

		if !prv.Mail.IsEmail.MatchString(rc.Email) {
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

		prv.Templates.ExecuteTemplate(w, "layout", vrr)

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
