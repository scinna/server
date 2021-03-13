package models

import "time"

type AuthToken struct {
	Token     string `db:"login_token"`
	LoginIP   string `db:"user_ip"`
	LastSeen  *time.Time `db:"last_seen"`
	CreatedAt *time.Time `db:"login_time"`
	RevokedAt *time.Time `db:"revoked_at"`
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
			LAST_SEEN   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			REVOKED_AT  TIMESTAMP DEFAULT NULL,
			PRIMARY KEY (LOGIN_TOKEN)
		);
	`
}
