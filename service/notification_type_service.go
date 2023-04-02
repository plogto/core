package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

func (s *Service) GetNotificationType(ctx context.Context, id uuid.UUID) (*db.NotificationType, error) {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByID(ctx, id)

	return notificationType, nil
}
