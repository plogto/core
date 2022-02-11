package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Notification struct {
	DB *pg.DB
}

func (p *Notification) GetNotificationsByReceiverIDAndPagination(receiverID string, limit, page int) (*model.Notifications, error) {
	var notifications []*model.Notification
	var offset = (page - 1) * limit

	query := p.DB.Model(&notifications).
		Where("receiver_id = ?", receiverID).
		Where("deleted_at is ?", nil).
		Order("created_at DESC").
		Returning("*")

	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	unreadNotificationsCount, _ := p.CountUnreadNotificationsByReceiverID(receiverID)

	return &model.Notifications{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		UnreadNotificationsCount: unreadNotificationsCount,
		Notifications:            notifications,
	}, err
}

func (p *Notification) CountUnreadNotificationsByReceiverID(receiverID string) (*int, error) {
	count, err := p.DB.Model((*model.Notification)(nil)).
		Where("receiver_id = ?", receiverID).
		Where("read = ?", false).
		Where("deleted_at is ?", nil).
		Returning("*").
		Count()

	return &count, err
}

func (n *Notification) CreateNotification(notification *model.Notification) (*model.Notification, error) {
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

func (n *Notification) RemoveNotification(notification *model.Notification) (*model.Notification, error) {
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

func (n *Notification) RemovePostNotifications(notification *model.Notification) (*model.Notification, error) {
	query := n.DB.Model(notification).
		Where("receiver_id = ?receiver_id").
		Where("post_id = ?post_id")

	_, err := query.Set("deleted_at = ?deleted_at").Returning("*").Update()
	return notification, err
}
