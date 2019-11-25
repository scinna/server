package model

// Picture represent one picture
type Picture struct {
	Model
	Title       string   `db:"title"`
	URLID       string   `db:"url_id"`
	Description string   `db:"descript"`
	Visibility  int8     `db:"visibility"`
	Creator     *AppUser `db:"creator" json:",omitempty"`
	Ext         string   `db:"ext" json:",omitempty"`
}
