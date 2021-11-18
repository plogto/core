package database

import (
	"fmt"

	"github.com/favecode/plog-core/graph/model"
	"github.com/go-pg/pg"
)

type NotificationType struct {
	DB *pg.DB
}

func (n *NotificationType) GetNotificationTypeByField(field, value string) (*model.NotificationType, error) {
	var notificationType model.NotificationType
	err := n.DB.Model(&notificationType).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(notificationType.ID) < 1 {
		return nil, nil
	}
	return &notificationType, err
}

func (n *NotificationType) GetNotificationTypeByID(id string) (*model.NotificationType, error) {
	return n.GetNotificationTypeByField("id", id)
}

func (n *NotificationType) GetNotificationTypeByName(name string) (*model.NotificationType, error) {
	return n.GetNotificationTypeByField("name", name)
}
