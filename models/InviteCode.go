package models

type InviteCode struct {
	InviteCode string `db:"invite_code"`
	InvitedBy  int    `db:"invited_by"`
	Used       bool   `db:"used"`
}
