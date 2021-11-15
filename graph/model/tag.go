package model

import (
	"time"
)

type Tag struct {
	tableName struct{}   `sql:"tag"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
