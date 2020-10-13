package dal

import (
	"github.com/scinna/server/models"
	"github.com/scinna/server/services"
)

// GetUserFromID returns a user from an id
func GetUserFromID(prv *services.Provider, id int) (*models.User, error){
	rq := `SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, VALIDATION_CODE FROM SCINNA_USER WHERE USER_ID = $1`
	row := prv.DB.QueryRowx(rq, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.StructScan(&user)

	return &user, err
}

// GetUserFromUsername returns a user from an id
func GetUserFromUsername(prv *services.Provider, username string) (*models.User, error){
	row := prv.DB.QueryRowx("SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, VALIDATION_CODE FROM SCINNA_USER WHERE USER_NAME = $1", username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.StructScan(&user)

	return &user, err
}

func Login(prv *services.Provider, user *models.User, ip string) (token string, err error) {
	row := prv.DB.QueryRow(`INSERT INTO LOGIN_TOKENS (USER_ID, USER_IP) VALUES ($1, $2) RETURNING LOGIN_TOKEN`, user.UserID, ip)
	if row.Err() != nil {
		return "", row.Err()
	}

	err = row.Scan(&token)
	return token, err
}

func FetchUserFromToken(prv *services.Provider, authToken string) (*models.User, error) {
	row := prv.DB.QueryRowx(`
		SELECT su.USER_ID, su.USER_NAME, su.USER_EMAIL, su.USER_PASSWORD, su.VALIDATED, su.VALIDATION_CODE
		FROM SCINNA_USER su
		INNER JOIN LOGIN_TOKENS lt ON lt.USER_ID = su.USER_ID 
		WHERE lt.LOGIN_TOKEN = $1
		  AND lt.REVOKED_AT IS NULL
	`, authToken)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.StructScan(&user)

	return &user, err
}