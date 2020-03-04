package auth

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/lib/pq"
	"github.com/scinna/server/dal"
	"github.com/scinna/server/model"
	"github.com/scinna/server/serrors"
	"github.com/scinna/server/services"
)

// GenerateToken create a token for the user passed in parameters
func GenerateToken(prv *services.Provider, ip string, u model.AppUser) (string, error) {
	rq := ` INSERT INTO LOGIN_TOKENS(ID_USR, IP) 
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
	// @TODO Rewrite everything in this >:(
	rq := ` SELECT ID_USR, REVOKED
			FROM LOGIN_TOKENS
			WHERE TOKEN = $1`

	var userID int
	var revoked bool
	row := prv.Db.QueryRow(rq, tokenStr)
	err := row.Scan(&userID, &revoked)
	if err != nil {

		if err == sql.ErrNoRows {
			// Happens when the token doesn't exists on the server
			return model.AppUser{}, serrors.ErrorTokenNotFound
		}

		errPost, ok := err.(*pq.Error)
		if ok && errPost.Code.Name() == serrors.PostgresError["InvalidUID"] {
			return model.AppUser{}, serrors.ErrorBadToken
		}
		return model.AppUser{}, err
	}

	if revoked {
		return model.AppUser{}, serrors.ErrorRevoked
	}

	return dal.GetUserByID(prv, userID)
}

// ValidateRequest retreives the token from a request, validate its token and return the corresponding user
func ValidateRequest(prv *services.Provider, w http.ResponseWriter, r *http.Request) (model.AppUser, error) {

	authToken, err := GetTokenFromRequest(r)
	if err != nil {
		return model.AppUser{}, err
	}

	return VerifyToken(prv, authToken)
}

// GetTokenFromRequest extracts the token from the request
func GetTokenFromRequest(r *http.Request) (string, error) {
	authToken := r.Header.Get("Authorization")

	if len(authToken) == 0 {
		return "", serrors.ErrorNoToken
	}

	splitToken := strings.Split(authToken, "Bearer ")
	if len(splitToken) > 1 {
		return splitToken[1], nil
	}

	return "", serrors.ErrorBadToken
}
