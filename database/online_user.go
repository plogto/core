package database

import (
	"github.com/favecode/plog-core/graph/model"
	"github.com/go-pg/pg"
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
