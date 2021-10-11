package model

import (
	"time"
)

type PostSave struct {
	tableName struct{}   `sql:"post_save"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
