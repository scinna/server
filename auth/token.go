package auth

import (
	"net/http"
	"strings"

	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// GenerateToken create a token for the user passed in parameters
func GenerateToken(prv *services.Provider, ip string, u model.AppUser) (string, error) {
	rq := ` INSERT INTO LOGIN_TOKENS(ID, IP) 
			VALUES ($1, $2)
			RETURNING TOKEN`

	var token string
	err := prv.Db.QueryRow(rq, u.ID, ip).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

// VerifyToken check if the token is valid. If so, it fetches the corresponding user from the database
func VerifyToken(prv *services.Provider, tokenStr string) (model.AppUser, error) {
	rq := ` SELECT ID, REVOKED
			FROM LOGIN_TOKENS
			WHERE TOKEN = $1`

	var userID int
	var revoked bool
	row := prv.Db.QueryRow(rq, tokenStr)
	err := row.Scan(&userID, &revoked)
	if err != nil {
		return model.AppUser{}, err
	}

	if revoked {
		return model.AppUser{}, serrors.ErrorRevoked
	}

	return dal.GetUserByID(prv, userID)
}

// ValidateRequest retreives the token from a request, validate its token and return the corresponding user
func ValidateRequest(prv *services.Provider, w http.ResponseWriter, r *http.Request) (model.AppUser, error) {
	authToken := r.Header.Get("Authorization")

	if len(authToken) == 0 {
		return model.AppUser{}, serrors.ErrorNoToken
	}

	splitToken := strings.Split(authToken, "Bearer ")
	if len(splitToken) > 1 {
		authToken = splitToken[1]
		return VerifyToken(prv, authToken)
	}
	return model.AppUser{}, serrors.ErrorBadToken
}
