package model

import (
	"fmt"
)

type Visibility int8

const (
	VISIBILITY_EVERYONE = 0
	VISIBILITY_UNLISTED = 1
	VISIBILITY_PRIVATE  = 2
)


type Picture struct {
	Model
	Title       string  `db:"title"`
	UrlId       string  `db:"url_id"`
	Description string  `db:"descript"`
	Visibility  int8    `db:"visibility"`
	Creator     *AppUser `db:"creator" json:",omitempty"`
}

/** For debug purposes **/
func (p *Picture) ToString() string {
	return fmt.Sprintf(`--- Picture %v ---
	ID: %v
	Title: %v
	Description: %v
	Created at: %v
	Visibility: %v
	Creator: %v
------------------`, p.Title, p.ID, p.Title, p.Description, p.CreatedAt, p.Visibility, p.Creator.Username)
}