package model

import (
	"time"
)

type Model struct {
	ID        int64      `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
}