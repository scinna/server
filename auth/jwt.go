package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/oxodao/scinna/dal"
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT create a token expiring in 20 minutes for the user passed in parameters
func GenerateJWT(prv *services.Provider, u model.AppUser) (string, error) {
	currTime := time.Now()
	tk := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": fmt.Sprintf("%v", u.ID), // For some reason strconv.Itoa doesn't compile @TODO
		"iss": "Scinna-Server",
		"exp": currTime.Add(20 * time.Minute).Unix(),
		"iat": currTime.Unix(),
	})

	token, err := tk.SignedString(prv.JwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// VerifyJWT check if the JWT is valid. If so, it fetches the corresponding user from the database
func VerifyJWT(prv *services.Provider, tokenStr string) (model.AppUser, error) {
	/**
	* We verify the JWT token standardly (Syntax OK + not expired)
	**/
	parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if ok := token.Method.Alg() == jwt.SigningMethodHS512.Alg(); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return prv.JwtKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return model.AppUser{}, serrors.ErrorBadToken
	}

	claims, _ := parsed.Claims.(jwt.MapClaims)
	a, _ := strconv.Atoi(claims["sub"].(string))
	u, err := dal.GetUserByID(prv, a)

	return u, err
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
		return VerifyJWT(prv, authToken)
	}
	return model.AppUser{}, serrors.ErrorBadToken
}
