package dal

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/log"
	"github.com/scinna/server/models"
	"github.com/scinna/server/requests"
)

type Registration struct {
	DB *sqlx.DB
}

func (r *Registration) FindInvite(invite string) *models.InviteCode {
	rq := `SELECT INVITE_CODE, INVITED_BY, USED FROM INVITE_CODE WHERE INVITE_CODE = $1`
	row := r.DB.QueryRowx(rq, invite)
	if row.Err() != nil {
		log.Warn(row.Err().Error())
		return nil
	}

	var ic models.InviteCode
	err := row.StructScan(&ic)
	if err != nil && err == sql.ErrNoRows {
		return nil
	}

	return &ic
}

func (r *Registration) DisableInvite(invite *models.InviteCode) {
	r.DB.Exec(`UPDATE invite_code SET used = true WHERE invite_code = $1`, invite.InviteCode)
}

func (r *Registration) RegisterUser(request *requests.RegisterRequest, canRegister bool) (string, error){
	rq := ` INSERT INTO SCINNA_USER (USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED)
			VALUES ($1, $2, $3, $4)
			RETURNING USER_ID, VALIDATION_CODE`

	row := r.DB.QueryRowx(rq, request.Username, request.Email, request.HashedPassword, canRegister)
	if row.Err() != nil {
		return "", row.Err()
	}

	var userid string
	var valcode string
	err := row.Scan(&userid, &valcode)
	if err != nil {
		return "", err
	}

	_, err = r.DB.Exec("INSERT INTO COLLECTIONS (title, user_id, visibility, default_collection) VALUES ($1, $2, $3, true)", "Default collection", userid, 0)

	return valcode, err
}

func (r *Registration) ValidateUser(validation string) string {
	var results []string
	err := r.DB.Select(&results, "UPDATE SCINNA_USER SET validated = true, validation_code = NULL WHERE validation_code = $1 RETURNING user_name", validation)
	if err != nil || len(results) == 0 {
		return ""
	}

	return results[0]
}

func (r *Registration) GenerateInviteIfNeeded(inviteCode string) (string, error) {
	row := r.DB.QueryRow("SELECT COUNT(*) FROM SCINNA_USER")
	if row.Err() != nil {
		return "", row.Err()
	}

	var amtUsers int
	err := row.Scan(&amtUsers)
	if err != nil {
		return "", row.Err()
	}

	if amtUsers > 0 {
		return "NONE", nil
	}

	rowx := r.DB.QueryRowx("SELECT INVITE_CODE, INVITED_BY, USED FROM INVITE_CODE WHERE USED = FALSE LIMIT 1")
	if rowx.Err() != nil {
		return "", rowx.Err()
	}

	var invite models.InviteCode
	err = rowx.StructScan(&invite)

	if err == sql.ErrNoRows {
		_, err = r.DB.Exec("INSERT INTO invite_code (invite_code, invited_by) VALUES ($1, -1)", inviteCode)
		return inviteCode, err
	}

	return invite.InviteCode, err
}