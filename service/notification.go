package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	onlineUserContext, err := middleware.GetCurrentOnlineUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	go func() {
		<-ctx.Done()
		s.mu.Lock()
		s.OnlineUser.DeleteOnlineUserBySocketId(onlineUserContext.SocketID)
		delete(s.Notifications, onlineUserContext.SocketID)
		s.mu.Unlock()
	}()

	s.mu.Lock()
	// // Keep a reference of the channel so that we can push changes into it when new messages are posted.
	notification := make(chan *model.Notification, 1)
	s.Notifications[onlineUserContext.SocketID] = notification
	s.mu.Unlock()

	return nil, nil
}
