package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/dto"
)

type Admin struct {
	DB *sqlx.DB
}

func (a Admin) ListUsers() ([]dto.AdminUser, error){
	users := []dto.AdminUser{}
	rq, err := a.DB.Queryx(`
		SELECT user_id, user_name, user_email, is_admin, validated, registered_at
		FROM scinna_user
	`)

	if err != nil {
		return users, err
	}

	for rq.Next() {
		u := dto.AdminUser{}
		_ = rq.StructScan(&u)
		users = append(users, u)
	}

	return users, nil
}