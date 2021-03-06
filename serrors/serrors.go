package serrors

import (
	"encoding/json"
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
func (s *SError) JSON() []byte {
	tx, err := json.Marshal(s)

	if err != nil {
		return []byte("{\"message\": \"Something went wrong encoding the error!\", \"errcode\": -1 }")
	}

	return tx
}

func (s *SError) Write(w http.ResponseWriter) {
	w.Header().Del("Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.HTTPError)
	_, _ = w.Write(s.JSON())
}

// WriteError writes the error to the response. Return false if there is no error
func WriteError(w http.ResponseWriter, err error) bool {
	if err != nil {
		cast, ok := err.(*SError)
		if ok {
			cast.Write(w)
			return true
		}

		NewUnknown(err).Write(w)
		return true
	}

	return false
}

// WriteLoggableError writes a generic error and save the real one to the database
func WriteLoggableError(w http.ResponseWriter, err error) bool {
	// @TODO Save the error, send a generic message with an error code to send to the admin
	if err != nil {
		cast, ok := err.(*SError)
		if ok {
			cast.Write(w)
			return true
		}

		NewUnknown(err).Write(w)
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
	@TODO: Standardize errors with the text being the translation key
	@TODO: Check if they are all required because there seems to be a lot of them here...
**/

// ErrorBadInviteCode is shown when the user tries to register with no invitation code or an invalid one
var ErrorBadInviteCode = New("The server requires a valid invitation code", 10, 400)

// ErrorInviteUsed is shown when the user tries to register with no invitation code or an invalid one
var ErrorInviteUsed = New("This invite code has already been used", 11, 400)

// ErrorUserExists is shown when the user tries to register with an already existing user
var ErrorUserExists = New("This username is already used", 12, 500)

// ErrorEmailExists is shown when the user tries to register with an already existing email
var ErrorEmailExists = New("This email is already used", 13, 500)

// ErrorInvalidRegistration is shown when the user tries to register with an empty username, password or email
var ErrorInvalidRegistration = New("Username, Email and Password can't be empty", 14, 400)

// InvalidUsernameOrPassword
var InvalidUsernameOrPassword = New("invalid_credentials", 15, 401)

// AccountNotValidated is sent when the user tries to login without being validated
var AccountNotValidated = New("This account is not validated", 16, 401)

// UserNeedValidationAdmin is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationAdmin = New("You have been registered. The admin now needs to validate your account", 120, 201)

// UserNeedValidationEmail is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationEmail = New("You have been registered. Please click the link on the email you have received to activate your account", 121, 201)

// UserRegistered is sent when the user was registered and the admin needs to validate the account
var UserRegistered = New("You can now use your account", 122, 201)

// InvalidType is sent when the user tries to upload a non autorized mime-type file
var InvalidType = New("Please upload a PICTURE file only", 4250, 400)

var InvalidValidationCode = New("This validation code does not exists or is already used", 4251, 400)

var NoToken = New("You can't access this resource without a valid authentication token", 401, 401)
var NotOwner = New("You don't have permission to reach this resource", 403, 403)


var CollectionNotFound = New("collection_not_found", 400, 404)
var CollectionAlreadyExists = New("collection_already_exists", 401, 400)
