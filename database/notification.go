package database

import (
	"github.com/favecode/plog-core/graph/model"
	"github.com/go-pg/pg/v10"
)

type Notification struct {
	DB *pg.DB
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

	if notification.CommentID != nil {
		query.Where("comment_id = ?comment_id")
	}

	_, err := query.Returning("*").SelectOrInsert()
	return notification, err
}
