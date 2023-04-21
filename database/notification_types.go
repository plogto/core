package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type NotificationTypes struct {
	Queries *db.Queries
}

func (n *NotificationTypes) GetNotificationTypeByID(ctx context.Context, id pgtype.UUID) (*db.NotificationType, error) {
	return n.Queries.GetNotificationTypeByID(ctx, id)
}

func (n *NotificationTypes) GetNotificationTypeByName(ctx context.Context, name db.NotificationTypeName) (*db.NotificationType, error) {
	return n.Queries.GetNotificationTypeByName(ctx, name)
}
