package middlewares

import (
	"context"
	"github.com/scinna/server/models"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
	"net/http"
	"strings"
)

func LoggedInMiddleware(prv *services.Provider) func(handler http.Handler) http.Handler {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, err := ValidateRequest(prv, r)
			if err != nil {
				serrors.NoToken.Write(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetTokenFromRequest extracts the token from the request
func GetTokenFromRequest(r *http.Request) (string, error) {
	authToken := r.Header.Get("Authorization")

	if len(authToken) == 0 {
		return "", serrors.NoToken
	}

	splitToken := strings.Split(authToken, "Bearer ")
	if len(splitToken) > 1 {
		return splitToken[1], nil
	}

	return "", serrors.NoToken
}

func ValidateRequest(prv *services.Provider, r *http.Request) (*models.User, error) {
	authToken, err := GetTokenFromRequest(r)
	if err != nil {
		return nil, err
	}
	return prv.Dal.User.FetchUserFromToken(authToken)
}
