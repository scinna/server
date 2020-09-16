package serrors

import (
	"encoding/json"
	"net/http"
)

// SError is a custom error type for Scinna
type SError struct {
	Message   string
	Errcode   int
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
	w.Write(s.JSON())
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
		Errcode:   errCode,
		HTTPError: err,
	}
}

// NewUnknown can be used to treat random errors
func NewUnknown(err error) *SError {
	return &SError{
		Message:   err.Error(),
		Errcode:   -1,
		HTTPError: http.StatusInternalServerError,
	}
}

// ErrorBadInviteCode is shown when the user tries to register with no invitation code or an invalid one
var ErrorBadInviteCode = New("The server requires a valid invitation code", 10, 400)

// ErrorUserExists is shown when the user tries to register with an already existing user
var ErrorUserExists = New("This username is already used", 11, 500)

// ErrorEmailExists is shown when the user tries to register with an already existing email
var ErrorEmailExists = New("This email is already used", 12, 500)

// ErrorInvalidRegistration is shown when the user tries to register with an empty username, password or email
var ErrorInvalidRegistration = New("Username, Email and Password can't be empty", 13, 400)

// InvalidUsernameOrPassword
var InvalidUsernameOrPassword = New("Invalid username or password", 14, 401)

// UserNeedValidationAdmin is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationAdmin = New("You have been registered. The admin now needs to validate your account", 120, 201)

// UserNeedValidationEmail is sent when the user was registered and the admin needs to validate the account
var UserNeedValidationEmail = New("You have been registered. Please click the link on the email you have received to activate your account", 121, 201)

// UserRegistered is sent when the user was registered and the admin needs to validate the account
var UserRegistered = New("You can now use your account", 122, 201)

var NoToken = New("You can't access this resource without a valid authentication token", 401, 401)