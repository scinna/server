package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/scinna/server/models"
)

type Server struct {
	DB *sqlx.DB
}

func (s *Server) ListInviteCode() {

}

func (s *Server) GenerateInviteCode(user *models.User, inviteCode string) error {
	_, err := s.DB.Exec("INSERT INTO INVITE_CODE (INVITE_CODE, INVITED_BY) VALUES ($1, $2)", inviteCode, user.UserID)
	return err
}
