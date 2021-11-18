package model

import (
	"time"
)

type NotificationType struct {
	tableName struct{}   `sql:"notification_type"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Template  string     `json:"template"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
