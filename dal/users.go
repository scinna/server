package dal

import (
	"github.com/oxodao/scinna/model"
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
