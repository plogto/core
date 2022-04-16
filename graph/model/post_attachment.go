package model

import (
	"time"
)

type PostAttachment struct {
	tableName struct{} `pg:"post_attachment"`
	ID        string
	PostID    string
	FileID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
