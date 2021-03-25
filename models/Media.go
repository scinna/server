package models

import (
	"bufio"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

	Path        string     `db:"path" json:"-"`
	Visibility  Visibility `db:"visibility"`
	PublishedAt time.Time  `db:"published_at"`
	Mimetype    string     `db:"mimetype" json:"-"`
	CustomData  CustomData `db:"custom_data" json:"custom_data"`
	Thumbnail   string     `db:"thumbnail" json:",omitempty"`

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

// GenerateThumbnail takes a source path for the picture and generates it's thumnail attribute. Only works for pictures for now
func (m *Media) GenerateThumbnail(source string) error {
	/**
	 * @TODO: Thumbnailise it
	 */
	f, err := os.Open(source)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	m.Thumbnail = fmt.Sprintf("data:%v;base64,%v", m.Mimetype, base64.StdEncoding.EncodeToString(content))

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
			CUSTOM_DATA  JSONB,
			THUMBNAIL    TEXT DEFAULT '',
			MIMETYPE     VARCHAR DEFAULT ''
		);
	`
}
