package models

import "time"

type AuthToken struct {
	Token    string
	LastSeen *time.Time
	Revoked  bool
}

func (at AuthToken) GetTableName() string {
	return "LOGIN_TOKENS"
}

func (at AuthToken) GenerateTable() string {
	return `
		CREATE TABLE LOGIN_TOKENS
		(
			LOGIN_TOKEN uuid      DEFAULT gen_random_uuid(),
			USER_ID     uuid REFERENCES SCINNA_USER (USER_ID) NOT NULL,
			USER_IP     VARCHAR                               NOT NULL,
			LOGIN_TIME  TIMESTAMP DEFAULT CURRENT_TIMESTAMP   NOT NULL,
			REVOKED_AT  TIMESTAMP DEFAULT NULL,
			PRIMARY KEY (LOGIN_TOKEN)
		);
	`
}
