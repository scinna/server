// Package serrors defines all the errors used in the software. This is so it doesn't conflict with the golang errors package
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

// ErrorTokenNotFound shows up whenever the request should be authed but the given token doesn't exists
var ErrorTokenNotFound *SError = New("This token doesn't exists", 399, http.StatusUnauthorized)

// ErrorNoToken shows up whenever the request should be authed and is not given any token
var ErrorNoToken *SError = New("No token found in the request", 400, http.StatusUnauthorized)

// ErrorBadToken shows up whenever the request should be authed but the given token is not in the correct format
var ErrorBadToken *SError = New("The given token is in the wrong format. It should be 'Bearer [TOKEN]'", 401, http.StatusUnauthorized)

// ErrorRevoked shows up whenever the request should be authed and the given token has been revoked
var ErrorRevoked *SError = New("This token has been revoked and can no longer be used", 403, http.StatusUnauthorized)

// ErrorPictureNotFound shows up whenever the picture supposedly removed doesn't exists from the DB (Should NEVER happens)
var ErrorPictureNotFound *SError = New("Picture not found", 404, http.StatusNotFound)

// ErrorBadRequest shows up whenever the client sends a malformated request (Can't read the body, bad JSON, etc...)
var ErrorBadRequest *SError = New("The server can't parse your request! (Are you sure all required fields were filled ?)", 405, http.StatusBadRequest)

// ErrorDatabase is just to throw when there is an error inserting/deleting but we don't care about it that much
var ErrorDatabase *SError = New("There was a database error!", 406, http.StatusInternalServerError)

// ErrorISE is just to throw when there is an error that is generic (Ex: can't seek to the beginning of a file)
var ErrorISE *SError = New("There was a server-side error", 407, http.StatusInternalServerError)

// ErrorGenerationUID is when the URL ID generator fails
var ErrorGenerationUID *SError = New("There was an error generating the unique ID", 408, http.StatusInternalServerError)

// ErrorWrongOwner is when the user request a picture or modify a picture that he doesn't own
var ErrorWrongOwner *SError = New("This picture doesn't belong to you", 409, http.StatusForbidden)

// ErrorBadFile happens when a user try to send a wrong file
var ErrorBadFile *SError = New("The file you are uploading is incorrect (Not an image or more than 10 meg)", 410, http.StatusBadRequest)

// ErrorSendingMail is thrown when the server failed to send an email
var ErrorSendingMail *SError = New("The server failed to send email. Please contact the administrator if needed", 411, http.StatusInternalServerError)

// ErrorMaxAttempts is thrown when the user tries too much to login
var ErrorMaxAttempts *SError = New("Calm down on login attempts!", 412, http.StatusBadRequest)

// ErrorUserNotFound is thrown when the user can't be found in the database
var ErrorUserNotFound *SError = New("User not found", 413, http.StatusBadRequest)

// ErrorInvalidCredentials is thrown when the user logs with a wrong password
var ErrorInvalidCredentials *SError = New("Invalid credentials", 414, http.StatusBadRequest)

// ErrorPrivatePicture is thrown when the user asks for a private picture that he doesn't own
var ErrorPrivatePicture *SError = New("This picture is private", 415, http.StatusForbidden)

// ErrorMissingURLID happens when the client request a route that requires a URL ID and doesn't feed it
var ErrorMissingURLID *SError = New("Request is missing the picture ID!", 418, http.StatusBadRequest)

// ErrorInvalidMimetype happens when you send a file that can't be uploaded to the server
var ErrorInvalidMimetype *SError = New("This file type can't be uploaded (Only jpeg, png or gif)", 419, http.StatusBadRequest)

// ErrorBadInviteCode happens when a user tries to register with an invalid invite code
var ErrorBadInviteCode *SError = New("Invalid invitation code", 420, http.StatusBadRequest)

// ErrorNotAdmin happens when a user to use an admin route without being one
var ErrorNotAdmin *SError = New("You are not an administrator", 421, http.StatusBadRequest)

/////// Registration errors

// ErrorRegDisabled gets thrown when a user tries to register while the registration are disabled
var ErrorRegDisabled *SError = New("Registration are disabled", 460, http.StatusBadRequest)

// ErrorRegExistingUser gets thrown when the user already exists
var ErrorRegExistingUser *SError = New("This username is already taken", 465, http.StatusConflict)

// ErrorRegExistingMail gets thrown when the user already exists
var ErrorRegExistingMail *SError = New("This email is already in use", 466, http.StatusConflict)

// ErrorRegBadUsername gets thrown when the user wants to register an invalid username (Either empty or blacklisted one)
var ErrorRegBadUsername *SError = New("This username is invalid (Either is empty or equals to 'me')", 467, http.StatusBadRequest)

// ErrorRegBadEmail gets thrown when the user wants to register an invalid username (Either empty or blacklisted one)
var ErrorRegBadEmail *SError = New("This email is invalid", 468, http.StatusBadRequest)

// ErrorAlreadyValidated shows up when you try to activate an already activated user
var ErrorAlreadyValidated *SError = New("This account is already activated", 469, http.StatusAlreadyReported)

// ErrorNoAccountValidation shows up when you try to activate a non existing validation token
var ErrorNoAccountValidation *SError = New("This activation token does not exists.", 470, http.StatusAlreadyReported)

// ErrorNotValidated shows up when the user account is not activated
var ErrorNotValidated *SError = New("This account is not activated", 471, http.StatusForbidden)

// ErrorRateLimited gets thrown when an IP requests more than X time
var ErrorRateLimited *SError = New("You've hit the rate limit for this API", 999, http.StatusBadRequest)
