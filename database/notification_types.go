package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type NotificationTypes struct {
	Queries *db.Queries
}

func (n *NotificationTypes) GetNotificationTypeByID(ctx context.Context, id uuid.UUID) (*db.NotificationType, error) {
	notificationType, _ := n.Queries.GetNotificationTypeByID(ctx, id)

	return notificationType, nil
}

func (n *NotificationTypes) GetNotificationTypeByName(ctx context.Context, name db.NotificationTypeName) (*db.NotificationType, error) {
	notificationType, _ := n.Queries.GetNotificationTypeByName(ctx, name)

	return notificationType, nil

}
