package models

type Media struct {
	MediaID string `db:"media_id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Path       string `db:"path" json:"-"`
	Visibility int    `db:"visibility"`
	Mimetype   string `db:"mimetype" json:"-"`

	User *User `db:"User"`
}

func (m Media) GetTableName() string {
	return "MEDIA"
}

func (m Media) GenerateTable() string {
	return `
		CREATE TABLE MEDIA
		(
			MEDIA_ID    VARCHAR(10) PRIMARY KEY,
			USER_ID     uuid REFERENCES SCINNA_USER (USER_ID) NOT NULL,
			CLC_ID      uuid REFERENCES COLLECTIONS (CLC_ID) NOT NULL,
			TITLE       VARCHAR,
			DESCRIPTION VARCHAR,
			PATH        VARCHAR,
			VISIBILITY  INTEGER,
			MIMETYPE    VARCHAR
		);
	`
}
