package dto

type ServerConfig struct {
	RegistrationAllowed bool
	Validation          string
	WebURL              string
	CustomBranding      string `json:",omitempty"`
	ScinnaVersion       string `json:",omitempty"`
}


type InviteCode struct {
	Code        string `db:"invite_code"`
	Author      string `db:"author"`
	GeneratedAt string `db:"generated_at"`
	Used        bool   `db:"used"`
}
