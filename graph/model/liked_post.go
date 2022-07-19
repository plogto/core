package model

import (
	"time"
)

type LikedPost struct {
	tableName struct{} `pg:"liked_posts"`
	ID        string
	UserID    string
	PostID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
