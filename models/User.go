package models

type User struct {
	UserID string `db:"user_id"`
	Name   string `db:"user_name"`
	Email  string `db:"user_email"`

	IsAdmin bool `db:"user_admin"`

	Password       string  `db:"user_password" json:"-"`
	Validated      bool    `db:"validated" json:"-"`
	ValidationCode *string `db:"validation_code" json:"-"`

	Collections []Collection
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
 
			USER_ADMIN      BOOL                 NOT NULL DEFAULT FALSE,

			INVITATION_CODE VARCHAR(10)                   DEFAULT NULL,

			VALIDATED       BOOL                 NOT NULL DEFAULT FALSE,
			VALIDATION_CODE VARCHAR              NULL     DEFAULT gen_random_uuid(),
			REGISTERED_AT   TIMESTAMP                     DEFAULT NOW()
		);
	`
}
