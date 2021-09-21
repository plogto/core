package model

import (
	"time"
)

type PostLike struct {
	tableName struct{}   `sql:"post_like"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
