package model

import (
	"time"
)

type Post struct {
	tableName struct{} `pg:"posts,discard_unknown_columns"`
	ID        string
	UserID    string
	ParentID  *string
	ChildID   *string
	Content   *string
	Url       string
	Status    PostStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
