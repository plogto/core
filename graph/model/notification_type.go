package model

import (
	"time"
)

type NotificationType struct {
	tableName struct{} `pg:"notification_types"`
	ID        string
	Name      string
	Template  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
