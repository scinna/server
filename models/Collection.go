package models

type Collection struct {
	Title      string  `db:"title"`
	User       *User   `db:"user"`
	Visibility int     `db:"visibility"`
	Medias     []Media `db:"-"`
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
			USER_ID            uuid REFERENCES SCINNA_USER (USER_ID) NOT NULL,
			VISIBILITY         INTEGER,
			DEFAULT_COLLECTION BOOL DEFAULT FALSE,
			UNIQUE (USER_ID, TITLE)
		);
	`
}
