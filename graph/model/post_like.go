package model

import (
	"time"
)

type PostLike struct {
	tableName struct{} `pg:"post_like"`
	ID        string
	UserID    string
	PostID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
