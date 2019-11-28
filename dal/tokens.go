package dal

import (
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// ListTokens returns the list of token for a user (Not the token itself, only other fields)
func ListTokens(prv *services.Provider, u model.AppUser) ([]model.LoginToken, error) {

	rq := `SELECT ID, CREATED_AT, IP, REVOKED FROM LOGIN_TOKENS WHERE ID_USR = $1`

	var tokens []model.LoginToken = []model.LoginToken{}
	rows, err := prv.Db.Queryx(rq, *u.ID)
	if err != nil {
		// Should never happen
		return []model.LoginToken{}, err
	}

	for rows.Next() {
		var t model.LoginToken
		rows.StructScan(&t)

		tokens = append(tokens, t)
	}

	return tokens, nil
}

// RevokeToken revokes an existing token
func RevokeToken(prv *services.Provider, idToken int, idUser int64) error {

	rq := `UPDATE LOGIN_TOKENS SET REVOKED = TRUE, REVOKED_DATE = CURRENT_TIMESTAMP WHERE ID = $1 AND ID_USR = $2`

	r, err := prv.Db.Exec(rq, idToken, idUser)
	if err != nil {
		return err
	}

	aff, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return serrors.ErrorTokenNotFound
	}

	return nil
}
