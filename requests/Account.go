package requests

type RegisterRequest struct {
	Username       string
	Email          string
	Password       string
	HashedPassword string `json:"-"`
	InviteCode     string
}

type LoginRequest struct {
	Username string
	Password string
}

type EditAccountRequest struct {
	Email           string
	CurrentPassword string
	Password        string
}
