package model

import (
	"fmt"
)

type AppUser struct {
	Model
	Email          string `db:"email"`
	Username       string `db:"username"`
	Password       string `db:"password"`
}

/** For debug purposes **/
func (a *AppUser) ToString() string {
	return fmt.Sprintf(`--- User %v ---
	ID: %v
	Username: %v
	Email: %v
	Created at: %v
------------------`, a.Username, a.ID, a.Username, a.Email, a.CreatedAt)
}