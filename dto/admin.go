package dto

type AdminUser struct {
	UserID string `db:"user_id"`
	Name   string `db:"user_name"`
	Email  string `db:"user_email"`

	IsAdmin bool `db:"is_admin"`

	Validated bool `db:"validated"`
}
