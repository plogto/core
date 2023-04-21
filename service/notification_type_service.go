package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

func (s *Service) GetNotificationType(ctx context.Context, id pgtype.UUID) (*db.NotificationType, error) {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByID(ctx, id)

	return notificationType, nil
}
