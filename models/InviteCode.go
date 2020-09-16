package models

type InviteCode struct {
	InviteID   int    `db:"invite_id"`
	InviteCode string `db:"invite_code"`
	InvitedBy  int    `db:"invited_by"`
	Used       bool   `db:"used"`
}
