package model

import (
	"time"
)

type Comment struct {
	tableName struct{}   `sql:"comment"`
	ID        string     `json:"id"`
	ParentID  *string    `json:"parent_id"`
	UserID    string     `json:"user_id"`
	PostID    string     `json:"post_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
