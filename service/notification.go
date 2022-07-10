package service

import (
	"context"
	"errors"
	"time"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

type CreateNotificationArgs struct {
	Name       string
	SenderID   string
	ReceiverID string
	Url        string
	PostID     *string
	ReplyID    *string
}

type RemovePostNotificationsArgs struct {
	ReceiverID string
	PostID     string
}

func (s *Service) GetNotifications(ctx context.Context, input *model.PaginationInput) (*model.Notifications, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	var limit int = constants.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	posts, _ := s.Notification.GetNotificationsByReceiverIDAndPagination(user.ID, limit, page)

	return posts, nil
}

func (s *Service) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	onlineUserContext, err := middleware.GetCurrentOnlineUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		s.OnlineUser.DeleteOnlineUserBySocketID(onlineUserContext.SocketID)
		delete(s.Notifications, onlineUserContext.SocketID)
		s.mu.Unlock()
	}()

	notification := make(chan *model.Notification, 1)
	s.mu.Lock()
	// Keep a reference of the channel so that we can push changes into it when new messages are posted.
	s.Notifications[onlineUserContext.SocketID] = notification
	s.mu.Unlock()

	return notification, nil
}

func (s *Service) CreateNotification(args CreateNotificationArgs) error {

	if args.SenderID != args.ReceiverID {
		notificationType, _ := s.NotificationType.GetNotificationTypeByName(args.Name)
		notification := &model.Notification{
			NotificationTypeID: notificationType.ID,
			SenderID:           args.SenderID,
			ReceiverID:         args.ReceiverID,
			URL:                args.Url,
			PostID:             args.PostID,
			ReplyID:            args.ReplyID,
		}

		s.Notification.CreateNotification(notification)

		onlineUser, _ := s.OnlineUser.GetOnlineUserByUserID(args.ReceiverID)

		if onlineUser != nil {
			s.mu.Lock()
			s.Notifications[onlineUser.SocketID] <- notification
			s.mu.Unlock()
		}
	}

	return nil
}

func (s *Service) RemoveNotification(args CreateNotificationArgs) error {
	notificationType, _ := s.NotificationType.GetNotificationTypeByName(args.Name)
	DeletedAt := time.Now()
	notification := &model.Notification{
		NotificationTypeID: notificationType.ID,
		SenderID:           args.SenderID,
		ReceiverID:         args.ReceiverID,
		PostID:             args.PostID,
		ReplyID:            args.ReplyID,
		DeletedAt:          &DeletedAt,
	}

	s.Notification.RemoveNotification(notification)

	// TODO: add removed type for Notification
	// onlineUser, _ := s.OnlineUser.GetOnlineUserByUserID(args.ReceiverID)

	// if onlineUser != nil {
	// 	s.mu.Lock()
	// 	s.Notifications[onlineUser.SocketID] <- notification
	// 	s.mu.Unlock()
	// }

	return nil
}

func (s *Service) RemoveNotifications(args RemovePostNotificationsArgs) error {
	DeletedAt := time.Now()
	notification := &model.Notification{
		ReceiverID: args.ReceiverID,
		PostID:     &args.PostID,
		DeletedAt:  &DeletedAt,
	}

	s.Notification.RemovePostNotifications(notification)

	return nil
}
