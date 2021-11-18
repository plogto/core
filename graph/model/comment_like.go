package model

import (
	"time"
)

type CommentLike struct {
	tableName struct{} `pg:"comment_like"`
	ID        string
	UserID    string
	CommentID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
