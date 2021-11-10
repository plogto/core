package model

import (
	"time"
)

type NotificationType struct {
	tableName struct{}   `sql:"notification_type"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Template  string     `json:"template"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
