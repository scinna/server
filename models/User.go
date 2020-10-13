package models

import "time"

type AuthToken struct {
	Token    string
	LastSeen *time.Time
	Revoked  bool
}

type User struct {
	UserID string `db:"user_id"`
	Name   string `db:"user_name"`
	Email  string `db:"user_email"`

	Password       string `db:"user_password" json:"-"`
	Validated      bool   `db:"validated" json:"-"`
	ValidationCode string `db:"validation_code" json:"-"`

	Medias []Media
}
