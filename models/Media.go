package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type MediaType = int

const (
	MEDIA_PICTURE  = 0
	MEDIA_VIDEO    = 1
	MEDIA_TXTBIN   = 2
	MEDIA_SHORTURL = 3
)

type Media struct {
	MediaID   string `db:"media_id"`
	MediaType int    `db:"media_type"`

	Title       string `db:"title" json:",omitempty"`
	Description string `db:"description" json:",omitempty"`

	Path          string     `db:"path" json:"-"`
	Visibility    Visibility `db:"visibility"`
	PublishedAt   time.Time  `db:"published_at"`
	ViewCount     int        `db:"view_count"`
	Mimetype      string     `db:"mimetype" json:"-"`
	CustomData    CustomData `db:"custom_data" json:"custom_data"`

	Collection *Collection `db:"collection" json:",omitempty"`

	User *User `db:"User" json:",omitempty"`
}

type CustomData map[string]interface{}

func (cd CustomData) Value() (driver.Value, error) {
	return json.Marshal(cd)
}

func (cd *CustomData) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("custom_data should be a byte array")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*cd, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("could not parse custom_data")
	}

	return nil
}

func (m Media) GetTableName() string {
	return "MEDIA"
}

func (m Media) GenerateTable() string {
	return `
		CREATE TABLE MEDIA
		(
			MEDIA_ID     VARCHAR(10) PRIMARY KEY,
			MEDIA_TYPE   INTEGER NOT NULL,
			USER_ID      uuid REFERENCES SCINNA_USER (USER_ID) ON DELETE CASCADE NOT NULL,
			CLC_ID       uuid REFERENCES COLLECTIONS (CLC_ID) ON DELETE CASCADE NOT NULL,
			TITLE        VARCHAR DEFAULT '',
			DESCRIPTION  VARCHAR DEFAULT '',
			PATH         VARCHAR DEFAULT '',
			VISIBILITY   INTEGER,
			VIEW_COUNT   INTEGER NOT NULL DEFAULT 0,
			PUBLISHED_AT TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			CUSTOM_DATA  JSONB,
			MIMETYPE     VARCHAR DEFAULT ''
		);
	`
}
