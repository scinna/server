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
