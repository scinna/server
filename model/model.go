package model

import "time"

// Model represents the standard attributes that a entity should have
type Model struct {
	ID        *int64     `db:"id" json:",omitempty"`
	CreatedAt *time.Time `db:"created_at" json:",omitempty"`
}
