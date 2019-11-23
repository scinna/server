package model

import (
	"fmt"
)

// Picture represent one picture
type Picture struct {
	Model
	Title       string   `db:"title"`
	URLID       string   `db:"url_id"`
	Description string   `db:"descript"`
	Visibility  int8     `db:"visibility"`
	Creator     *AppUser `db:"creator" json:",omitempty"`
}

// ToString is a debug method to print one picture, should be removed
func (p *Picture) ToString() string {
	return fmt.Sprintf(`--- Picture %v ---
	ID: %v
	URLID: %v
	Title: %v
	Description: %v
	Created at: %v
	Visibility: %v
	Creator: %v
------------------`, p.Title, p.ID, p.URLID, p.Title, p.Description, p.CreatedAt, p.Visibility, p.Creator.Username)
}
