package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Notifications struct {
	Queries *db.Queries
}

func (n *Notifications) CreateNotification(ctx context.Context, arg db.CreateNotificationParams) (*db.Notification, error) {
	notification, _ := n.Queries.GetNotification(ctx, db.GetNotificationParams{
		NotificationTypeID: arg.NotificationTypeID,
		SenderID:           arg.SenderID,
		ReceiverID:         arg.ReceiverID,
		PostID:             arg.PostID,
		ReplyID:            arg.ReplyID,
		Url:                arg.Url,
	})

	if notification != nil {
		return notification, nil
	}

	return n.Queries.CreateNotification(ctx, arg)
}

func (n *Notifications) GetNotificationByID(ctx context.Context, id pgtype.UUID) (*db.Notification, error) {
	return n.Queries.GetNotificationByID(ctx, id)
}

func (n *Notifications) GetNotificationsByReceiverIDAndPageInfo(ctx context.Context, receiverID pgtype.UUID, limit int32, after time.Time) (*model.Notifications, error) {
	var edges []*model.NotificationsEdge
	var endCursor string

	notifications, _ := n.Queries.GetNotificationsByReceiverIDAndPageInfo(ctx, db.GetNotificationsByReceiverIDAndPageInfoParams{
		Limit:      limit,
		ReceiverID: receiverID,
		CreatedAt:  after,
	})

	totalCount, _ := n.Queries.CountNotificationsByReceiverIDAndPageInfo(ctx, db.CountNotificationsByReceiverIDAndPageInfoParams{
		ReceiverID: receiverID,
		CreatedAt:  after,
	})

	unreadNotificationsCount, _ := n.CountUnreadNotificationsByReceiverID(ctx, receiverID)

	for _, value := range notifications {
		edges = append(edges, &model.NotificationsEdge{Node: &db.Notification{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Notifications{
		TotalCount:               totalCount,
		Edges:                    edges,
		UnreadNotificationsCount: unreadNotificationsCount,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (n *Notifications) CountUnreadNotificationsByReceiverID(ctx context.Context, receiverID pgtype.UUID) (int64, error) {
	count, _ := n.Queries.CountUnreadNotificationsByReceiverID(ctx, receiverID)

	return count, nil
}

func (n *Notifications) UpdateReadNotifications(ctx context.Context, receiverID pgtype.UUID) (bool, error) {
	_, err := n.Queries.UpdateReadNotifications(ctx, receiverID)

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (n *Notifications) RemoveNotification(ctx context.Context, arg db.RemoveNotificationParams) (*db.Notification, error) {
	return n.Queries.RemoveNotification(ctx, arg)
}

func (n *Notifications) RemovePostNotificationsByPostID(ctx context.Context, postID pgtype.UUID) ([]*db.Notification, error) {
	DeletedAt := time.Now()

	notifications, _ := n.Queries.RemovePostNotificationsByPostID(ctx, db.RemovePostNotificationsByPostIDParams{
		PostID:    postID,
		DeletedAt: &DeletedAt,
	})

	return notifications, nil
}
