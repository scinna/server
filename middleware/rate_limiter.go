package middleware

import (
	"net/http"

	"github.com/scinna/server/dal"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
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
