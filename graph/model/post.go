package model

import (
	"time"
)

type Post struct {
	tableName  struct{} `pg:"post,discard_unknown_columns"`
	ID         string
	UserID     string
	Content    string
	Url        string
	Attachment *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `pg:"-,soft_delete"`
}
