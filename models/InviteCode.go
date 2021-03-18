package models

import "time"

type InviteCode struct {
	InviteCode  string    `db:"invite_code"`
	InvitedBy   *string   `db:"invited_by"`
	GeneratedAt time.Time `db:"generated_at"`
	Used        bool      `db:"used"`
}

func (ic InviteCode) GetTableName() string {
	return "INVITE_CODE"
}

func (ic InviteCode) GenerateTable() string {
	return `
		CREATE TABLE INVITE_CODE
		(
			INVITE_CODE  VARCHAR(10) PRIMARY KEY NOT NULL,
			INVITED_BY   uuid NULL,
			GENERATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			USED         BOOL DEFAULT FALSE
		);
	`
}
