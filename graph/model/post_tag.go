package model

import (
	"time"
)

type PostTag struct {
	tableName struct{} `pg:"post_tag"`
	ID        string
	TagID     string
	PostID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
