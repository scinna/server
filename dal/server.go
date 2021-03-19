package dal

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/dto"
	"github.com/scinna/server/models"
)

type Server struct {
	DB *sqlx.DB
}

func (s *Server) ListInviteCode() ([]dto.InviteCode, error) {
	rows, err := s.DB.Queryx(`
		SELECT INVITE_CODE, USER_NAME AS AUTHOR, GENERATED_AT, USED
		FROM INVITE_CODE
			 INNER JOIN SCINNA_USER ON INVITED_BY = USER_ID
		UNION
		SELECT INVITE_CODE, 'Server' AS AUTHOR, GENERATED_AT, USED
		FROM INVITE_CODE
		WHERE INVITED_BY IS NULL 
		ORDER BY GENERATED_AT DESC
`)

	if err != nil {
		return nil, err
	}

	// Do not replace, we need a empty array and not nil
	invites := []dto.InviteCode{}
	for rows.Next() {
		currInvite := dto.InviteCode{}
		err = rows.StructScan(&currInvite)
		if err != nil {
			break
		}

		invites = append(invites, currInvite)
	}

	return invites, err
}

func (s *Server) GenerateInviteCode(user *models.User, inviteCode string) error {
	_, err := s.DB.Exec("INSERT INTO INVITE_CODE (INVITE_CODE, INVITED_BY) VALUES ($1, $2)", inviteCode, user.UserID)
	return err
}

func (s *Server) Delete(code string) error {
	result, err := s.DB.Exec("DELETE FROM INVITE_CODE WHERE INVITE_CODE = $1", code)
	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if i != 1 {
		// @TODO: translate blabla
		return errors.New("not found")
	}

	return nil
}
