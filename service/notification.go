package service

import (
	"context"
	"errors"
	"time"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
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

func (s *Service) GetNotifications(ctx context.Context, input *model.PageInfoInput) (*model.Notifications, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	pageInfoInput := util.ExtractPageInfo(input)

	return s.Notifications.GetNotificationsByReceiverIDAndPageInfo(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}

func (s *Service) GetNotificationByID(ctx context.Context, id string) (*model.Notification, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	return s.Notifications.GetNotificationByID(id)
}

func (s *Service) GetNotification(ctx context.Context) (<-chan *model.NotificationsEdge, error) {
	onlineUserContext, err := middleware.GetCurrentOnlineUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		s.OnlineUsers.DeleteOnlineUserBySocketID(onlineUserContext.SocketID)
		delete(s.OnlineNotifications, onlineUserContext.SocketID)
		s.mu.Unlock()
	}()

	notificationEdge := make(chan *model.NotificationsEdge, 1)
	s.mu.Lock()
	// Keep a reference of the channel so that we can push changes into it when new messages are posted.
	s.OnlineNotifications[onlineUserContext.SocketID] = notificationEdge
	s.mu.Unlock()

	return notificationEdge, nil
}

func (s *Service) CreateNotification(args CreateNotificationArgs) error {

	if args.SenderID != args.ReceiverID {
		notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(args.Name)
		notification := &model.Notification{
			NotificationTypeID: notificationType.ID,
			SenderID:           args.SenderID,
			ReceiverID:         args.ReceiverID,
			URL:                args.Url,
			PostID:             args.PostID,
			ReplyID:            args.ReplyID,
		}

		notificationEdge := &model.NotificationsEdge{
			Cursor: util.ConvertCreateAtToCursor(*notification.CreatedAt),
			Node:   notification,
		}

		s.Notifications.CreateNotification(notification)

		onlineUser, _ := s.OnlineUsers.GetOnlineUserByUserID(args.ReceiverID)

		if onlineUser != nil {
			s.mu.Lock()
			s.OnlineNotifications[onlineUser.SocketID] <- notificationEdge
			s.mu.Unlock()
		}
	}

	return nil
}

func (s *Service) RemoveNotification(args CreateNotificationArgs) error {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(args.Name)
	DeletedAt := time.Now()
	notification := &model.Notification{
		NotificationTypeID: notificationType.ID,
		SenderID:           args.SenderID,
		ReceiverID:         args.ReceiverID,
		PostID:             args.PostID,
		ReplyID:            args.ReplyID,
		DeletedAt:          &DeletedAt,
	}

	s.Notifications.RemoveNotification(notification)

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

	s.Notifications.RemovePostNotifications(notification)

	return nil
}
