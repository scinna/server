package dal

import (
	"github.com/oxodao/scinna/services"
	"github.com/oxodao/scinna/model"
)

func GetUser(p *services.Provider, username string) (model.AppUser, error) {
	rq := ` SELECT ID, CREATED_AT, EMAIL, USERNAME, PASSWORD
			FROM APPUSER
			WHERE USERNAME = $1`

	var user model.AppUser
	err := p.Db.QueryRowx(rq, username).StructScan(&user)

	return user, err
}