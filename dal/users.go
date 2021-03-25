package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
	"time"
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
	rq := `SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED, VALIDATION_CODE, IS_ADMIN, REGISTERED_AT FROM SCINNA_USER WHERE USER_ID = $1`
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
	row := u.DB.QueryRowx("SELECT USER_ID, USER_NAME, USER_EMAIL, USER_PASSWORD, IS_ADMIN, VALIDATED, VALIDATION_CODE, REGISTERED_AT FROM SCINNA_USER WHERE USER_NAME = $1", username)
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
		SELECT su.USER_ID, su.USER_NAME, su.USER_EMAIL, su.USER_PASSWORD, su.VALIDATED, su.VALIDATION_CODE, su.IS_ADMIN, su.REGISTERED_AT
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

func (u *User) TokenBelongsToUser(user *models.User, token string) bool {
	rq := u.DB.QueryRow("SELECT TRUE FROM LOGIN_TOKENS WHERE USER_ID = $1 AND LOGIN_TOKEN = $2", user.UserID, token)
	if rq.Err() != nil {
		return false
	}

	var belongsToUser bool
	err := rq.Scan(&belongsToUser)
	if err != nil || !belongsToUser {
		return false
	}

	return true
}

func (u *User) RevokeToken(authToken string) (*time.Time, error){
	rows := u.DB.QueryRowx("UPDATE LOGIN_TOKENS SET REVOKED_AT = NOW() WHERE LOGIN_TOKEN = $1 RETURNING REVOKED_AT", authToken)
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var revokedAt time.Time
	err := rows.Scan(&revokedAt)

	return &revokedAt, err
}

func (u *User) FetchUserTokens(user *models.User) ([]models.AuthToken, error){
	var tokens []models.AuthToken
	// We take all non revoked token + revoked tokens revoked less than a month ago
	rows, err := u.DB.Queryx(`
		SELECT LOGIN_TOKEN, USER_IP, LOGIN_TIME, LAST_SEEN, REVOKED_AT
		FROM LOGIN_TOKENS
		WHERE USER_ID = $1
		  AND (REVOKED_AT = NULL OR LOGIN_TIME > (NOW() - interval '1 month'))
		ORDER BY LOGIN_TIME DESC
	`, user.UserID)

	if err != nil {
		return tokens, err
	}

	for rows.Next() {
		token := models.AuthToken{}
		err = rows.StructScan(&token)
		if err != nil {
			return tokens, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (u *User) UpdatePassword(user *models.User, hashedPassword string) {
	_, _ = u.DB.Exec("UPDATE SCINNA_USER SET USER_PASSWORD = $2 WHERE USER_ID = $1", user.UserID, hashedPassword)
}

func (u *User) UpdateEmail(user *models.User, email string, autovalidate bool) {
	if autovalidate {
		_, _ = u.DB.Exec("UPDATE SCINNA_USER SET USER_EMAIL = $2 WHERE USER_ID = $1", user.UserID, email)
	} else {
		_, _ = u.DB.Exec("UPDATE SCINNA_USER SET USER_EMAIL = $2, VALIDATED = false WHERE USER_ID = $1", user.UserID, email)
	}
}