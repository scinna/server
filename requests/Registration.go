package requests

type RegisterRequest struct {
	Username   string
	Email      string
	Password   string
	InviteCode string
}

type LoginRequest struct {
	Username string
	Password string
}