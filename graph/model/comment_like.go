package model

import (
	"time"
)

type CommentLike struct {
	tableName struct{}   `sql:"comment_like"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	CommentID string     `json:"comment_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
