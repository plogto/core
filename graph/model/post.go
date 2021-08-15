package model

import (
	"time"
)

type Post struct {
	tableName struct{}   `sql:"post"`
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	Content   string     `json:"content"`
	Status    *int       `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
