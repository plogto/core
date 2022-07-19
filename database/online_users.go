package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type OnlineUsers struct {
	DB *pg.DB
}

func (o *OnlineUsers) CreateOnlineUser(onlineUser *model.OnlineUser) (*model.OnlineUser, error) {
	_, err := o.DB.Model(onlineUser).Returning("*").Insert()
	return onlineUser, err
}

func (o *OnlineUsers) DeleteOnlineUserBySocketID(socketID string) (*model.OnlineUser, error) {
	var onlineUser = &model.OnlineUser{
		SocketID: socketID,
	}
	_, err := o.DB.Model(onlineUser).Where("socket_id = ?socket_id").Where("deleted_at is ?", nil).Returning("*").Delete()
	return onlineUser, err
}

func (o *OnlineUsers) DeleteAllOnlineUsers() ([]*model.OnlineUser, error) {
	var onlineUsers []*model.OnlineUser

	_, err := o.DB.Model(&onlineUsers).Where("true").Returning("*").Delete()
	return onlineUsers, err
}

func (o *OnlineUsers) GetOnlineUserByUserID(userID string) (*model.OnlineUser, error) {
	var onlineUser model.OnlineUser
	err := o.DB.Model(&onlineUser).Where("user_id = ?", userID).Where("deleted_at is ?", nil).First()
	if len(onlineUser.ID) < 1 {
		return nil, nil
	}
	return &onlineUser, err
}
