package model

// AppUser represents one user in database
type AppUser struct {
	Model
	Email     string `db:"email" json:",omitempty"`
	Username  string `db:"username"`
	Password  string `db:"password" json:"-"`
	Validated bool   `db:"validated" json:"-"` // Not needed to serialize it since all request that returns a user it must be validated
}
