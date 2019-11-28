package model

// UserRoleUser returns the number associated with the user role
const UserRoleUser int = 0

// UserRoleAdmin returns the number associated with the admin role
const UserRoleAdmin int = 100

// AppUser represents one user in database
type AppUser struct {
	Model
	Email           string `db:"email" json:",omitempty"`
	Role            int    `db:"role"`
	Username        string `db:"username"`
	Password        string `db:"password" json:"-"`
	Validated       bool   `db:"validated" json:"-"` // Not needed to serialize it since all request that returns a user it must be validated
	ValidationToken string `db:"validation_token" json:"-"`
}
