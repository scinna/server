package models

type Collection struct {
	Title      string  `db:"title"`
	User       *User   `db:"user"`
	Visibility int     `db:"visibility"`
	Medias     []Media `db:"-"`
}
