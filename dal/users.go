package dal

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/routes"
	"github.com/oxodao/scinna/services"
)

// GetUser fetches one user from the database given its username
func GetUser(p *services.Provider, username string) (model.AppUser, error) {
	rq := ` SELECT ID, CREATED_AT, EMAIL, USERNAME, PASSWORD
			FROM APPUSER
			WHERE USERNAME = $1`

	var user model.AppUser
	err := p.Db.QueryRowx(rq, username).StructScan(&user)

	return user, err
}

// GetUserByID fetches one user from the database given its id
func GetUserByID(p *services.Provider, id int) (model.AppUser, error) {
	rq := ` SELECT ID, CREATED_AT, EMAIL, USERNAME, PASSWORD
			FROM APPUSER
			WHERE ID = $1`

	var user model.AppUser
	err := p.Db.QueryRowx(rq, id).StructScan(&user)

	return user, err
}

// ReachedMaxAttempts returns whether the user has tried to login more than 10 times in the last five minutes
func ReachedMaxAttempts(prv *services.Provider, u model.AppUser) (bool, error) {
	rq := `
			SELECT 
				CASE WHEN (CREATED_AT > CURRENT_TIMESTAMP - INTERVAL '5 min') THEN 
					CASE WHEN ((SELECT COUNT(*) FROM LOGIN_ATTEMPT WHERE ID = $1 AND CREATED_AT > CURRENT_TIMESTAMP - INTERVAL '5 min') >= $2) THEN
						TRUE
					ELSE
						FALSE
					END
				ELSE 
					FALSE 
				END 
			FROM LOGIN_ATTEMPT
			WHERE ID = $1
			ORDER BY CREATED_AT DESC
			LIMIT 1`

	var canLogIn bool
	err := prv.Db.Get(&canLogIn, rq, u.ID, 10)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		fmt.Println(err)
		return true, err
	}

	return canLogIn, nil
}

// InsertFailedLoginAttempt inserts in the database a entry to log when a user fails his login attempt
func InsertFailedLoginAttempt(prv *services.Provider, u model.AppUser, ip string) error {
	rq := ` INSERT INTO LOGIN_ATTEMPT(ID, IP) 
			VALUES ($1, $2)`
	_, err := prv.Db.Exec(rq, u.ID, strings.Split(ip, ":")[0])
	return err
}

// RegisterUser inserts a non validated user in the DB
func RegisterUser(prv *services.Provider, rc *routes.RegisterRequest) error {

	hPass, err := prv.HashPassword(rc.Password)
	if err != nil {
		return err
	}

	rq := ` INSERT INTO APPUSER(USERNAME, EMAIL, PASSWORD) 
			VALUES ($1, $2, $3)`
	_, err = prv.Db.Exec(rq, rc.Username, rc.Email, hPass)

	return err

}
