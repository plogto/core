package model

import (
	"time"
)

type PostLike struct {
	tableName struct{}   `sql:"post_like"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
