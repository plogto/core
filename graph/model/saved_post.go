package model

import (
	"time"
)

type SavedPost struct {
	tableName struct{} `pg:"saved_posts"`
	ID        string
	UserID    string
	PostID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
