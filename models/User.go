package models

import "time"

type User struct {
	UserID string `db:"user_id"`
	Name   string `db:"user_name"`
	Email  string `db:"user_email"`

	IsAdmin bool `db:"is_admin"`

	Password       string  `db:"user_password" json:"-"`
	Validated      bool    `db:"validated" json:"-"`
	ValidationCode *string `db:"validation_code" json:"-"`
	ResetPasswordCode *string `db:"reset_pwd_code" json:"-"`

	RegisteredAt *time.Time `db:"registered_at"`
}

func (u User) GetTableName() string {
	return "SCINNA_USER"
}

func (u User) GenerateTable() string {
	return `
		CREATE TABLE SCINNA_USER
		(
			USER_ID         uuid PRIMARY KEY              default gen_random_uuid(),
			USER_NAME       VARCHAR(30)  UNIQUE  NOT NULL,
			USER_EMAIL      VARCHAR(255) UNIQUE  NOT NULL,
			USER_PASSWORD   VARCHAR              NOT NULL,
 
			IS_ADMIN        BOOL                 NOT NULL DEFAULT FALSE,

			INVITATION_CODE VARCHAR(10)                   DEFAULT NULL,

			VALIDATED       BOOL                 NOT NULL DEFAULT FALSE,
			VALIDATION_CODE VARCHAR              NULL     DEFAULT gen_random_uuid(),
			RESET_PWD_CODE  VARCHAR              NULL,
			REGISTERED_AT   TIMESTAMP                     DEFAULT NOW()
		);
	`
}
