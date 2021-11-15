package model

import (
	"time"
)

type Post struct {
	tableName  struct{}   `sql:"post" pg:",discard_unknown_columns"`
	ID         string     `json:"id"`
	UserID     string     `json:"userId"`
	Content    string     `json:"content"`
	Url        string     `json:"url"`
	Attachment string     `json:"attachment"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" sql:",soft_delete"`
}
