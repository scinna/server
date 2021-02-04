package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type User struct {
	DB *sqlx.DB
}

func (u *User) InsertUser(user *models.User) error {
	rq := `INSERT INTO SCINNA_USER (USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, IS_ADMIN) VALUES ($1, $2, $3, $4, $5) RETURNING USER_ID`
	row := u.DB.QueryRowx(rq, user.Name, user.Email, user.Password, user.Validated, user.IsAdmin)
	if row.Err() != nil {
		return row.Err()
	}

	return row.StructScan(user)
}

// GetUserFromID returns a user from an id
func (u *User) GetUserFromID(id int) (*models.User, error){
	rq := `SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, VALIDATION_CODE, IS_ADMIN FROM SCINNA_USER WHERE USER_ID = $1`
	row := u.DB.QueryRowx(rq, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.StructScan(&user)

	return &user, err
}

// GetUserFromUsername returns a user from an id
func (u *User) GetUserFromUsername(username string) (*models.User, error){
	row := u.DB.QueryRowx("SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, VALIDATION_CODE FROM SCINNA_USER WHERE USER_NAME = $1", username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.StructScan(&user)

	return &user, err
}

func (u *User) Login(user *models.User, ip string) (token string, err error) {
	row := u.DB.QueryRow(`INSERT INTO LOGIN_TOKENS (USER_ID, USER_IP) VALUES ($1, $2) RETURNING LOGIN_TOKEN`, user.UserID, ip)
	if row.Err() != nil {
		return "", row.Err()
	}

	err = row.Scan(&token)
	return token, err
}

func (u *User) FetchUserFromToken(authToken string) (*models.User, error) {
	row := u.DB.QueryRowx(`
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
