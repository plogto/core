package model

import "time"

type NotificationType struct {
	tableName struct{} `pg:"notification_types"`
	ID        string
	Name      string
	Template  string
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
