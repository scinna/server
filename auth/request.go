package auth

import (
	"net/http"

	"github.com/oxodao/scinna/serrors"
)

// RespondError will return true if there is an error
func RespondError(w http.ResponseWriter, err error) bool {
	if err != nil {
		if err == serrors.ErrorNoToken || err == serrors.ErrorBadToken {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
	return err != nil
}
