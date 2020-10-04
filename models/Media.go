package models

type Media struct {
	MediaID int `db:"media_id"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Path string `db:"path"`

	Visibility int `db:"visibility"`
}
