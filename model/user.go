package model

import (
	"fmt"
)

// AppUser represents one user in database
type AppUser struct {
	Model
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password" json:"-"`
}

// ToString is a debug method to print one picture, should be removed
func (a *AppUser) ToString() string {
	return fmt.Sprintf(`--- User %v ---
	ID: %v
	Username: %v
	Email: %v
	Created at: %v
------------------`, a.Username, a.ID, a.Username, a.Email, a.CreatedAt)
}
