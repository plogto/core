package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type NotificationTypes struct {
	Queries *db.Queries
}

func (n *NotificationTypes) GetNotificationTypeByID(ctx context.Context, id string) (*db.NotificationType, error) {
	// FIXME
	ID, _ := uuid.Parse(id)
	notificationType, err := n.Queries.GetNotificationTypeByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return notificationType, nil
}

func (n *NotificationTypes) GetNotificationTypeByName(ctx context.Context, name db.NotificationTypeName) (*db.NotificationType, error) {
	notificationType, err := n.Queries.GetNotificationTypeByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return notificationType, nil

}
