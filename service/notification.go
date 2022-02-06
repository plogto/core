package service

import (
	"context"
	"errors"
	"time"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

type CreateNotificationArgs struct {
	Name       string
	SenderId   string
	ReceiverId string
	Url        string
	PostId     *string
	ReplyId    *string
}

type RemovePostNotificationsArgs struct {
	ReceiverId string
	PostId     string
}

func (s *Service) GetNotifications(ctx context.Context, input *model.PaginationInput) (*model.Notifications, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	posts, _ := s.Notification.GetNotificationsByReceiverIdAndPagination(user.ID, limit, page)

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
		s.OnlineUser.DeleteOnlineUserBySocketId(onlineUserContext.SocketID)
		delete(s.Notifications, onlineUserContext.SocketID)
		s.mu.Unlock()
	}()

	notification := make(chan *model.Notification, 1)
	s.mu.Lock()
	// // Keep a reference of the channel so that we can push changes into it when new messages are posted.
	s.Notifications[onlineUserContext.SocketID] = notification
	s.mu.Unlock()

	return notification, nil
}

func (s *Service) CreateNotification(args CreateNotificationArgs) error {

	if args.SenderId != args.ReceiverId {
		notificationType, _ := s.NotificationType.GetNotificationTypeByName(args.Name)
		notification := &model.Notification{
			NotificationTypeID: notificationType.ID,
			SenderID:           args.SenderId,
			ReceiverID:         args.ReceiverId,
			URL:                args.Url,
			PostID:             args.PostId,
			ReplyID:            args.ReplyId,
		}

		s.Notification.CreateNotification(notification)

		onlineUser, _ := s.OnlineUser.GetOnlineUserByUserId(args.ReceiverId)

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
		SenderID:           args.SenderId,
		ReceiverID:         args.ReceiverId,
		PostID:             args.PostId,
		ReplyID:            args.ReplyId,
		DeletedAt:          &DeletedAt,
	}

	s.Notification.RemoveNotification(notification)

	// TODO: add revmoved type for Notification
	// onlineUser, _ := s.OnlineUser.GetOnlineUserByUserId(args.ReceiverId)

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
		ReceiverID: args.ReceiverId,
		PostID:     &args.PostId,
		DeletedAt:  &DeletedAt,
	}

	s.Notification.RemovePostNotifications(notification)

	return nil
}
