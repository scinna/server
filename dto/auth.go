package dto

import "github.com/scinna/server/models"

type AuthDto struct {
	*models.User
	Token string
}
