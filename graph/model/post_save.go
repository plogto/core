package model

import (
	"time"
)

type PostSave struct {
	tableName struct{} `pg:"post_save"`
	ID        string
	UserID    string
	PostID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
