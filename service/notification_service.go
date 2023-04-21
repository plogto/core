package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateNotificationArgs struct {
	Name       db.NotificationTypeName
	SenderID   pgtype.UUID
	ReceiverID pgtype.UUID
	Url        string
	PostID     pgtype.UUID
	ReplyID    pgtype.UUID
}

type RemovePostNotificationsArgs struct {
	ReceiverID string
	PostID     string
}

type CreatePostMentionNotificationsArgs struct {
	UserIDs  []pgtype.UUID
	Post     model.Post
	SenderID pgtype.UUID
}

func (s *Service) GetNotifications(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Notifications, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	pagination := util.ExtractPageInfo(pageInfo)

	return s.Notifications.GetNotificationsByReceiverIDAndPageInfo(ctx, user.ID, pagination.First, pagination.After)
}

func (s *Service) GetNotificationByID(ctx context.Context, id pgtype.UUID) (*db.Notification, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, err
	}

	return s.Notifications.GetNotificationByID(ctx, id)
}

func (s *Service) GetNotification(ctx context.Context) (*model.NotificationsEdge, error) {
	// TODO: rewrite this function

	return nil, nil
}

func (s *Service) CreateNotification(ctx context.Context, args CreateNotificationArgs) error {

	if args.SenderID != args.ReceiverID {
		notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(ctx, args.Name)

		if notificationType != nil {
			s.Notifications.CreateNotification(ctx, db.CreateNotificationParams{
				NotificationTypeID: notificationType.ID,
				SenderID:           args.SenderID,
				ReceiverID:         args.ReceiverID,
				Url:                args.Url,
				PostID:             args.PostID,
				ReplyID:            args.ReplyID,
			})
		}

		// TODO: handle online users

	}

	return nil
}

func (s *Service) RemoveNotification(ctx context.Context, args CreateNotificationArgs) error {
	notificationType, _ := s.NotificationTypes.GetNotificationTypeByName(ctx, args.Name)
	DeletedAt := time.Now()
	s.Notifications.RemoveNotification(ctx, db.RemoveNotificationParams{
		NotificationTypeID: notificationType.ID,
		SenderID:           args.SenderID,
		ReceiverID:         args.ReceiverID,
		PostID:             args.PostID,
		ReplyID:            args.ReplyID,
		DeletedAt:          &DeletedAt,
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
			s.CreateNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameMentionInPost,
				SenderID:   args.SenderID,
				ReceiverID: receiverID,
				Url:        "/p/" + args.Post.Url,
				PostID:     args.Post.ID,
			})
		}
	}
}
