package models

type InviteCode struct {
	InviteCode string `db:"invite_code"`
	InvitedBy  int    `db:"invited_by"`
	Used       bool   `db:"used"`
}

func (ic InviteCode) GetTableName() string {
	return "INVITE_CODE"
}

func (ic InviteCode) GenerateTable() string {
	return `
		CREATE TABLE INVITE_CODE
		(
			INVITE_CODE VARCHAR(10) PRIMARY KEY NOT NULL,
			INVITED_BY  INTEGER,
			USED        BOOL DEFAULT FALSE
		);
	`
}
