package service

import (
	"context"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) AddOnlineUser(ctx context.Context, onlineUserContext *middleware.OnlineUserContext) {
	onlineUser := &model.OnlineUser{
		UserID:    onlineUserContext.User.ID,
		SocketID:  onlineUserContext.SocketID,
		Token:     onlineUserContext.Token,
		UserAgent: onlineUserContext.UserAgent,
	}

	s.OnlineUser.CreateOnlineUser(onlineUser)
}
