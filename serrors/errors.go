// Package serrors defines all the errors used in the software. This is so it doesn't conflict with the golang errors package
package serrors

import (
	"errors"
)

// ErrorNoToken shows up whenever the request should be authed and is not given any token
var ErrorNoToken error = errors.New("No token found in the request")

// ErrorBadToken shows up whenever the request should be authed but the given token is not in the correct format
var ErrorBadToken error = errors.New("The given token is in the wrong format. It should be 'Bearer [JWT TOKEN]'")

// ErrorRevoked shows up whenever the request should be authed and the given token has been revoked
var ErrorRevoked error = errors.New("This token has been revoked and can no longer be used")

// ErrorPictureNotFound shows up whenever the picture supposedly removed doesn't exists from the DB (Should NEVER happens)
var ErrorPictureNotFound error = errors.New("The picture can't be removed from the DB since it doesn't exists")
