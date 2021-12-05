package service

import (
	"context"

	"github.com/favecode/plog-core/graph/model"
)

func (s *Service) GetNotificationType(ctx context.Context, id string) (*model.NotificationType, error) {
	notificationType, _ := s.NotificationType.GetNotificationTypeByID(id)

	return notificationType, nil
}
