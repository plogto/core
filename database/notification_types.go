package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type NotificationTypes struct {
	Queries *db.Queries
}

func (n *NotificationTypes) GetNotificationTypeByID(ctx context.Context, id uuid.UUID) (*db.NotificationType, error) {
	return util.HandleDBResponse(n.Queries.GetNotificationTypeByID(ctx, id))
}

func (n *NotificationTypes) GetNotificationTypeByName(ctx context.Context, name db.NotificationTypeName) (*db.NotificationType, error) {
	return util.HandleDBResponse(n.Queries.GetNotificationTypeByName(ctx, name))
}
