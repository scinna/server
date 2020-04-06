package model

// Media represent one media (Picture or video)
type Media struct {
	Model
	Title       string   `db:"title"`
	URLID       string   `db:"url_id"`
	Description string   `db:"descript"`
	Visibility  int8     `db:"visibility"`
	Creator     *AppUser `db:"creator" json:",omitempty"`
	Thumbnail   string   `db:"thumbnail" json:"thumbnail,omitempty"`
	Ext         string   `db:"ext" json:",omitempty"`
}
