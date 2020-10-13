package dal

import (
	"database/sql"
	"github.com/scinna/server/log"
	"github.com/scinna/server/models"
	"github.com/scinna/server/requests"
	"github.com/scinna/server/services"
)

func FindInvite(prv *services.Provider, invite string) *models.InviteCode {
	rq := `SELECT INVITE_CODE, INVITED_BY, USED FROM INVITE_CODE WHERE INVITE_CODE = $1`
	row := prv.DB.QueryRowx(rq, invite)
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

func DisableInvite(prv *services.Provider, invite *models.InviteCode) {
	prv.DB.Exec(`UPDATE invite_code SET used = true WHERE invite_code = $1`, invite.InviteCode)
}

func RegisterUser(prv *services.Provider, request *requests.RegisterRequest) (string, error){
	rq := ` INSERT INTO SCINNA_USER (USER_NAME, USER_EMAIL, USER_PASSWORD, VALIDATED)
			VALUES ($1, $2, $3, $4)
			RETURNING VALIDATION_CODE`

	pwd, err := prv.HashPassword(request.Password)
	if err != nil {
		return "", err
	}

	row := prv.DB.QueryRowx(rq, request.Username, request.Email, pwd, prv.Config.Registration.Validation == "open")
	if row.Err() != nil {
		return "", row.Err()
	}

	var valcode string
	err = row.Scan(&valcode)

	return valcode, err
}

func ValidateUser(prv *services.Provider, validation string) string {
	var results []string
	err := prv.DB.Select(&results, "UPDATE SCINNA_USER SET validated = true, validation_code = NULL WHERE validation_code = $1 RETURNING user_name", validation)
	if err != nil || len(results) == 0 {
		return ""
	}

	return results[0]
}

func GenerateInviteIfNeeded(prv *services.Provider) (string, error) {
	row := prv.DB.QueryRow("SELECT COUNT(*) FROM SCINNA_USER")
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

	rowx := prv.DB.QueryRowx("SELECT INVITE_CODE, INVITED_BY, USED FROM INVITE_CODE WHERE USED = FALSE LIMIT 1")
	if rowx.Err() != nil {
		return "", rowx.Err()
	}

	var invite models.InviteCode
	err = rowx.StructScan(&invite)

	if err == sql.ErrNoRows {
		inviteCode, err := prv.GenerateUID()
		if err != nil {
			return "", err
		}

		_, err = prv.DB.Exec("INSERT INTO invite_code (invite_code, invited_by) VALUES ($1, -1)", inviteCode)
		return inviteCode, err
	}

	return invite.InviteCode, err
}