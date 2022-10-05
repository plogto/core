package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type NotificationTypes struct {
	DB *pg.DB
}

func (n *NotificationTypes) GetNotificationTypeByField(field, value string) (*model.NotificationType, error) {
	var notificationType model.NotificationType
	err := n.DB.Model(&notificationType).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(notificationType.ID) < 1 {
		return nil, nil
	}
	return &notificationType, err
}

func (n *NotificationTypes) GetNotificationTypeByID(id string) (*model.NotificationType, error) {
	return n.GetNotificationTypeByField("id", id)
}

func (n *NotificationTypes) GetNotificationTypeByName(name model.NotificationTypeName) (*model.NotificationType, error) {
	return n.GetNotificationTypeByField("name", string(name))
}
