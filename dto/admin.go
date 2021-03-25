package dto

import "time"

type AdminUser struct {
	UserID string `db:"user_id"`
	Name   string `db:"user_name"`
	Email  string `db:"user_email"`

	IsAdmin bool `db:"is_admin"`

	Validated    bool       `db:"validated"`
	RegisteredAt *time.Time `db:"registered_at"`
}
