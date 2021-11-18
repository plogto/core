package model

import (
	"time"
)

type Connection struct {
	tableName   struct{} `pg:"connection"`
	ID          string
	FollowingID string
	FollowerID  string
	Status      *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `pg:"-,soft_delete"`
}
