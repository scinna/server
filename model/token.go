package model

import "time"

// LoginToken is the token that authenticate the user
type LoginToken struct {
	Model
	IDUser    int       `db:"id_usr" json:",omitempty"`
	Token     string    `db:"token" json:",omitempty"`
	IP        string    `db:"ip"`
	Revoked   bool      `db:"revoked"`
	RevokedAt time.Time `db:"revoked_at"`
}
