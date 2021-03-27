package serrors

import (
	"encoding/json"
	"github.com/scinna/server/translations"
	"net/http"
)

// SError is a custom error type for Scinna
type SError struct {
	Message   string
	ErrCode   int
	HTTPError int `json:"-"`
}

func (s SError) Error() string {
	return s.Message
}

// JSON return the error as a JSON to be sent to the client
func (s *SError) JSON(r *http.Request) []byte {
	// @TODO: If not a custom error, do not translate it
	translatedError := *s
	translatedError.Message = translations.T(r, "errors." + s.Message)
	tx, err := json.Marshal(translatedError)

	if err != nil {
		return []byte("{\"message\": \"Something went wrong encoding the error!\", \"errcode\": -1 }")
	}

	return tx
}

func (s *SError) Write(w http.ResponseWriter, r *http.Request) {
	w.Header().Del("Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.HTTPError)
	_, _ = w.Write(s.JSON(r))
}

// WriteError writes the error to the response. Return false if there is no error
func WriteError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		cast, ok := err.(*SError)
		if ok {
			cast.Write(w, r)
			return true
		}

		NewUnknown(err).Write(w, r)
		return true
	}

	return false
}

// New creates a new error
func New(msg string, errCode, err int) *SError {
	return &SError{
		Message:   msg,
		ErrCode:   errCode,
		HTTPError: err,
	}
}

// NewUnknown can be used to treat random errors
func NewUnknown(err error) *SError {
	return &SError{
		Message:   err.Error(),
		ErrCode:   -1,
		HTTPError: http.StatusInternalServerError,
	}
}

/**
	@TODO: Check if they are all required because there seems to be a lot of them here...
**/

// ErrorBadInviteCode is shown when the user tries to register with no invitation code or an invalid one
var ErrorBadInviteCode = New("invite.bad_code", 10, 400)

// ErrorInviteUsed is shown when the user tries to register with no invitation code or an invalid one
var ErrorInviteUsed = New("invite.used", 11, 400)

// InvalidEmail is shown when the registration email is invalid
var InvalidEmail = New("registration.invalid_email", 17, 500)

// ErrorInvalidRegistration is shown when the user tries to register with an empty username, password or email
var ErrorInvalidRegistration = New("registration.invalid_form", 14, 400)

// InvalidUsernameOrPassword
var InvalidUsernameOrPassword = New("login.invalid_credentials", 15, 401)

// AccountNotValidated is sent when the user tries to login without being validated
var AccountNotValidated = New("login.not_validated", 16, 401)

// UserNeedValidationAdmin is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationAdmin = New("registration.validation.admin", 120, 201)

// UserNeedValidationEmail is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationEmail = New("registration.validation.email", 121, 201)

// UserRegistered is sent when the user was registered and the admin needs to validate the account
var UserRegistered = New("registration.validation.none", 122, 201)

// InvalidType is sent when the user tries to upload a non autorized mime-type file
var InvalidType = New("invalid_filetype", 4250, 400)

var InvalidValidationCode = New("registration.validation.unknown", 4251, 400)

var NoToken = New("authorization.no_token", 401, 401)
var NotOwner = New("authorization.not_owner", 403, 403)

var CollectionNotFound = New("collection.not_found", 400, 404)
var CollectionAlreadyExists = New("collection.already_exists", 401, 400)


var MediaNoThumbnail = New("media.no_thumbnail", 405, 400)