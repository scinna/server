package model

import "time"

// InvitationCode represent an invitation code
type InvitationCode struct {
	ID          int64     `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	GeneratedBy AppUser   `db:"creator"`
	Code        string    `db:"code"`
}
