package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type OnlineUser struct {
	DB *pg.DB
}

func (o *OnlineUser) CreateOnlineUser(onlineUser *model.OnlineUser) (*model.OnlineUser, error) {
	_, err := o.DB.Model(onlineUser).Returning("*").Insert()
	return onlineUser, err
}

func (o *OnlineUser) DeleteOnlineUserBySocketId(socketId string) (*model.OnlineUser, error) {
	var onlineUser = &model.OnlineUser{
		SocketID: socketId,
	}
	_, err := o.DB.Model(onlineUser).Where("socket_id = ?socket_id").Where("deleted_at is ?", nil).Returning("*").Delete()
	return onlineUser, err
}

func (o *OnlineUser) GetOnlineUserByUserId(userId string) (*model.OnlineUser, error) {
	var onlineUser model.OnlineUser
	err := o.DB.Model(&onlineUser).Where("user_id = ?", userId).Where("deleted_at is ?", nil).First()
	if len(onlineUser.ID) < 1 {
		return nil, nil
	}
	return &onlineUser, err
}
