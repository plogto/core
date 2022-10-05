package model

import "time"

type NotificationType struct {
	tableName struct{} `pg:"notification_types"`
	ID        string
	Name      NotificationTypeName
	Template  string
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
