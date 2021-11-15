package model

import (
	"time"
)

type PostSave struct {
	tableName struct{}   `sql:"post_save"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	PostID    string     `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
