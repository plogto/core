package model

import (
	"time"
)

type PostTag struct {
	tableName struct{}   `sql:"post_tag"`
	ID        string     `json:"id"`
	TagID     string     `json:"tag_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
