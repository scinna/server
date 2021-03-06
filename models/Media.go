package models

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Media struct {
	MediaID string `db:"media_id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Path        string     `db:"path" json:"-"`
	Visibility  Visibility `db:"visibility"`
	PublishedAt time.Time  `db:"published_at"`
	Mimetype    string     `db:"mimetype" json:"-"`
	Thumbnail   string     `db:"thumbnail"`

	Collection *Collection `db:"collection" json:",omitempty"`

	User *User `db:"User" json:",omitempty"`
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
			MEDIA_ID    VARCHAR(10) PRIMARY KEY,
			USER_ID     uuid REFERENCES SCINNA_USER (USER_ID) ON DELETE CASCADE NOT NULL,
			CLC_ID      uuid REFERENCES COLLECTIONS (CLC_ID) ON DELETE CASCADE NOT NULL ,
			TITLE       VARCHAR,
			DESCRIPTION VARCHAR,
			PATH        VARCHAR,
			VISIBILITY  INTEGER,
			THUMBNAIL   TEXT,
			MIMETYPE    VARCHAR
		);
	`
}
