package model

import (
	"time"
)

type Comment struct {
	tableName struct{} `pg:"comment"`
	ID        string
	ParentID  *string
	UserID    string
	PostID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
