package model

import (
	"time"
)

type PostMention struct {
	tableName struct{} `pg:"post_mentions"`
	ID        string
	UserID    string
	PostID    string
	CreatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
