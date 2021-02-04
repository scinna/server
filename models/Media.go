package models

import "time"

type Media struct {
	MediaID string `db:"media_id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Path        string    `db:"path" json:"-"`
	Visibility  int       `db:"visibility"`
	PublishedAt time.Time `db:"published_at"`
	Mimetype    string    `db:"mimetype" json:"-"`

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
			USER_ID     uuid REFERENCES SCINNA_USER (USER_ID) ON DELETE CASCADE NOT NULL,
			CLC_ID      uuid REFERENCES COLLECTIONS (CLC_ID) ON DELETE CASCADE NOT NULL ,
			TITLE       VARCHAR,
			DESCRIPTION VARCHAR,
			PATH        VARCHAR,
			VISIBILITY  INTEGER,
			MIMETYPE    VARCHAR
		);
	`
}
