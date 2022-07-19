package service

import (
	"context"

	"github.com/plogto/core/graph/model"
)

func (s *Service) GetNotificationType(ctx context.Context, id string) (*model.NotificationType, error) {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByID(id)

	return notificationType, nil
}
