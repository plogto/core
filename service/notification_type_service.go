package service

import (
	"context"

	"github.com/plogto/core/db"
)

func (s *Service) GetNotificationType(ctx context.Context, id string) (*db.NotificationType, error) {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByID(ctx, id)

	return notificationType, nil
}
