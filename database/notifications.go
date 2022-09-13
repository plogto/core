package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Notifications struct {
	DB *pg.DB
}

func (n *Notifications) GetNotificationByID(id string) (*model.Notification, error) {
	var notification model.Notification
	err := n.DB.Model(&notification).Where("id = ?", id).Where("deleted_at is ?", nil).First()
	return &notification, err
}

func (n *Notifications) GetNotificationsByReceiverIDAndPageInfo(receiverID string, limit int, after string) (*model.Notifications, error) {
	var notifications []*model.Notification
	var edges []*model.NotificationsEdge
	var endCursor string

	query := n.DB.Model(&notifications).
		Where("receiver_id = ?", receiverID).
		Where("deleted_at is ?", nil).
		Order("created_at DESC")

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	unreadNotificationsCount, _ := n.CountUnreadNotificationsByReceiverID(receiverID)

	for _, value := range notifications {
		edges = append(edges, &model.NotificationsEdge{Node: &model.Notification{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > limit {
		hasNextPage = true
	}

	return &model.Notifications{
		TotalCount:               &totalCount,
		Edges:                    edges,
		UnreadNotificationsCount: unreadNotificationsCount,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (n *Notifications) CountUnreadNotificationsByReceiverID(receiverID string) (*int, error) {
	count, err := n.DB.Model((*model.Notification)(nil)).
		Where("receiver_id = ?", receiverID).
		Where("read = ?", false).
		Where("deleted_at is ?", nil).
		Returning("*").
		Count()

	return &count, err
}

func (n *Notifications) CreateNotification(notification *model.Notification) (*model.Notification, error) {
	query := n.DB.Model(notification).
		Where("notification_type_id = ?notification_type_id").
		Where("sender_id = ?sender_id").
		Where("receiver_id = ?receiver_id").
		Where("deleted_at is ?", nil)

	if notification.PostID != nil {
		query.Where("post_id = ?post_id")
	}

	if notification.ReplyID != nil {
		query.Where("reply_id = ?reply_id")
	}

	_, err := query.Returning("*").SelectOrInsert()
	return notification, err
}

func (n *Notifications) RemoveNotification(notification *model.Notification) (*model.Notification, error) {
	query := n.DB.Model(notification).
		Where("notification_type_id = ?notification_type_id").
		Where("sender_id = ?sender_id").
		Where("receiver_id = ?receiver_id")

	if notification.PostID != nil {
		query.Where("post_id = ?post_id")
	}

	if notification.ReplyID != nil {
		query.Where("reply_id = ?reply_id")
	}

	_, err := query.Set("deleted_at = ?deleted_at").Returning("*").Update()
	return notification, err
}

func (n *Notifications) RemovePostNotifications(notification *model.Notification) (*model.Notification, error) {
	query := n.DB.Model(notification).
		Where("receiver_id = ?receiver_id").
		Where("post_id = ?post_id")

	_, err := query.Set("deleted_at = ?deleted_at").Returning("*").Update()
	return notification, err
}

func (n *Notifications) UpdateReadNotifications(receiverID string) (bool, error) {
	var notifications []*model.Notification

	query := n.DB.Model(&notifications).
		Where("receiver_id = ?", receiverID)

	_, err := query.Set("read = ?", true).Returning("*").Update()

	if err != nil {
		return false, nil
	}

	return true, nil
}
