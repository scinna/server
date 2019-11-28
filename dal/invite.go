package dal

import (
	"github.com/oxodao/scinna/model"
	"github.com/oxodao/scinna/serrors"
	"github.com/oxodao/scinna/services"
)

// GetInviteCode returns the invite code for the given
func GetInviteCode(prv *services.Provider, inviteCode string) (model.InvitationCode, error) {
	rq := `
		SELECT ic.ID, ic.CREATED_AT, ic.CODE,
			   au.ID AS "creator.id", au.CREATED_AT AS "creator.created_at", au.ROLE as "creator.role", au.EMAIL as "creator.email", au.USERNAME AS "creator.username"
		FROM INVITATION_CODE ic
			 INNER JOIN APPUSER au ON au.ID = ic.GENERATED_BY
		WHERE CODE = $1`

	var invcode model.InvitationCode
	err := prv.Db.QueryRowx(rq, inviteCode).StructScan(&invcode)
	if err != nil {
		return model.InvitationCode{}, serrors.ErrorBadInviteCode
	}

	return invcode, nil
}

// GenerateInviteCode create a one-time use invite code
func GenerateInviteCode(prv *services.Provider, u *model.AppUser) (string, error) {
	rq := `
		INSERT INTO INVITATION_CODE (GENERATED_BY, CODE)
		VALUES ($1, $2)`

	inviteCode, err := prv.GenerateUID()
	if err != nil {
		return "", err
	}

	_, err = prv.Db.Exec(rq, *u.ID, inviteCode)

	return inviteCode, err
}

// RevokeInviteCode removes the invite code when used by an user. It returns whether the invite code existed in the first place
func RevokeInviteCode(prv *services.Provider, inviteCode string) error {
	rq := `
		DELETE FROM INVITATION_CODE
		WHERE CODE = $1`

	_, err := prv.Db.Exec(rq, inviteCode)
	if err != nil {
		return serrors.ErrorBadInviteCode
	}

	return nil
}
