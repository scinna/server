/** Scinna Errors, just to not conflict **/
package serrors

import (
	"errors"
)

var NoTokenError error = errors.New("No token found in the request")
var BadTokenError error = errors.New("The given token is in the wrong format. It should be 'Bearer [JWT TOKEN]'")

