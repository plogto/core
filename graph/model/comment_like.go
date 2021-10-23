package model

import (
	"time"
)

type CommentLike struct {
	tableName struct{}   `sql:"comment_like"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	CommentID string     `json:"comment_id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
