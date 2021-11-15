package model

import (
	"time"
)

type PostTag struct {
	tableName struct{}   `sql:"post_tag"`
	ID        string     `json:"id"`
	TagID     string     `json:"tag_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
