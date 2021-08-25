package model

import (
	"time"
)

type Follower struct {
	tableName  struct{}   `sql:"follower"`
	ID         string     `json:"id"`
	UserID     string     `json:"user_id"`
	FollowerID string     `json:"follower_id"`
	Status     *int       `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"-" sql:",soft_delete"`
}
