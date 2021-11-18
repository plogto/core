package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) GetNotificationType(ctx context.Context, id string) (*model.NotificationType, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	notificationType, _ := s.NotificationType.GetNotificationTypeByID(id)

	return notificationType, nil
}
