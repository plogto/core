package model

import (
	"time"
)

type Connection struct {
	tableName   struct{}   `sql:"connection"`
	ID          string     `json:"id"`
	FollowingID string     `json:"following_id"`
	FollowerID  string     `json:"follower_id"`
	Status      *int       `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-" sql:",soft_delete"`
}
