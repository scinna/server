package middleware

import (
	"net/http"

	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// RateLimiterMiddlewareFunc limits the calls to the API
func RateLimiterMiddlewareFunc(prv *services.Provider, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if serrors.WriteError(w, dal.APIRateLimiting(prv, r)) {
			return
		}

		next.ServeHTTP(w, r)
	})
}
