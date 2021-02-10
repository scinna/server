package models

type Collection struct {
	CollectionID string       `db:"clc_id"`
	Title        string       `db:"title"`
	User         *User        `db:"user" json:",omitempty"`
	Visibility   int          `db:"visibility"`
	IsDefault    bool         `db:"default_collection"`
	Collections  []Collection `db:"collections"`
	Medias       []Media      `db:"-" json:",omitempty"`
}

func (c Collection) GetTableName() string {
	return "COLLECTIONS"
}

func (c Collection) GenerateTable() string {
	return `
		CREATE TABLE COLLECTIONS
		(
			CLC_ID             uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			TITLE              VARCHAR,
			USER_ID            uuid REFERENCES SCINNA_USER (USER_ID) ON DELETE CASCADE NOT NULL,
			VISIBILITY         INTEGER,
			DEFAULT_COLLECTION BOOL DEFAULT FALSE,
			UNIQUE (USER_ID, TITLE)
		);
	`
}
