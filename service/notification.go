package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateNotificationArgs struct {
	Name       db.NotificationTypeName
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Url        string
	PostID     uuid.NullUUID
	ReplyID    uuid.NullUUID
}

type RemovePostNotificationsArgs struct {
	ReceiverID string
	PostID     string
}

type CreatePostMentionNotificationsArgs struct {
	UserIDs  []string
	Post     model.Post
	SenderID string
}

func (s *Service) GetNotifications(ctx context.Context, input *model.PageInfoInput) (*model.Notifications, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	pageInfoInput := util.ExtractPageInfo(input)

	return s.Notifications.GetNotificationsByReceiverIDAndPageInfo(ctx, user.ID, int32(pageInfoInput.First), pageInfoInput.After)
}

func (s *Service) GetNotificationByID(ctx context.Context, id string) (*db.Notification, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	return s.Notifications.GetNotificationByID(ctx, id)
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

func (s *Service) CreateNotification(ctx context.Context, args CreateNotificationArgs) error {

	if args.SenderID != args.ReceiverID {
		notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(ctx, args.Name)

		notification, _ := s.Notifications.CreateNotification(ctx, db.CreateNotificationParams{
			NotificationTypeID: notificationType.ID,
			SenderID:           args.SenderID,
			ReceiverID:         args.ReceiverID,
			Url:                args.Url,
			PostID:             args.PostID,
			ReplyID:            args.ReplyID,
		})

		notificationEdge := &model.NotificationsEdge{
			Cursor: util.ConvertCreateAtToCursor(notification.CreatedAt),
			Node:   notification,
		}

		onlineUser, _ := s.OnlineUsers.GetOnlineUserByUserID(args.ReceiverID.String())

		if onlineUser != nil {
			s.mu.Lock()
			s.OnlineNotifications[onlineUser.SocketID] <- notificationEdge
			s.mu.Unlock()
		}
	}

	return nil
}

func (s *Service) RemoveNotification(ctx context.Context, args CreateNotificationArgs) error {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(ctx, args.Name)
	DeletedAt := sql.NullTime{time.Now(), true}
	s.Notifications.RemoveNotification(ctx, db.RemoveNotificationParams{
		NotificationTypeID: notificationType.ID,
		SenderID:           args.SenderID,
		ReceiverID:         args.ReceiverID,
		PostID:             args.PostID,
		ReplyID:            args.ReplyID,
		DeletedAt:          DeletedAt,
	})

	// TODO: add removed type for Notification
	// onlineUser, _ := s.OnlineUser.GetOnlineUserByUserID(args.ReceiverID)

	// if onlineUser != nil {
	// 	s.mu.Lock()
	// 	s.Notifications[onlineUser.SocketID] <- notification
	// 	s.mu.Unlock()
	// }

	return nil
}

func (s *Service) ReadNotifications(ctx context.Context) (*bool, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	status, _ := s.Notifications.UpdateReadNotifications(ctx, user.ID)

	return &status, nil
}

func (s *Service) CreatePostMentionNotifications(ctx context.Context, args CreatePostMentionNotificationsArgs) {
	for _, receiverID := range args.UserIDs {
		if receiverID != args.SenderID {
			// FIXME
			senderID, _ := uuid.Parse(args.SenderID)
			receiverID, _ := uuid.Parse(receiverID)
			postID, _ := uuid.Parse(args.Post.ID)
			s.CreateNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameMentionInPost,
				SenderID:   senderID,
				ReceiverID: receiverID,
				Url:        "/p/" + args.Post.Url,
				PostID:     uuid.NullUUID{postID, true},
			})
		}
	}
}
