package models

type User struct {
	UserID         int    `db:"user_id"`
	Name           string `db:"user_name"`
	Email          string `db:"user_email"`

	Password       string `db:"user_password" json:"-"`
	Validated      bool   `db:"validated" json:"-"`
	ValidationCode string `db:"validation_code" json:"-"`
}
