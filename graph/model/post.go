package model

import (
	"time"
)

type Post struct {
	tableName struct{}   `sql:"post" pg:",discard_unknown_columns"`
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	Content   string     `json:"content"`
	Url       string     `json:"url"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
