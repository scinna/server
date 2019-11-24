package model

// LoginToken is the token that authenticate the user
type LoginToken struct {
	Model
	Token string `db:"token"`
	IP    string `db:"ip"`
}
