package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
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
