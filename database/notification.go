package database

import (
	"github.com/favecode/plog-core/graph/model"
	"github.com/go-pg/pg/v10"
)

type Notification struct {
	DB *pg.DB
}

func (n *Notification) CreateNotification(notification *model.Notification) (*model.Notification, error) {
	_, err := n.DB.Model(notification).Returning("*").Insert()
	return notification, err
}
