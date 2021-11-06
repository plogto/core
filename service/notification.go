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
		s.mu.Unlock()
	}()
	// r.mu.Lock()
	// // Keep a reference of the channel so that we can push changes into it when new messages are posted.
	// r.ChatObservers[id] = msgs
	// r.mu.Unlock()
	// // This is optional, and this allows newly subscribed clients to get a list of all the messages that have been
	// // posted so far. Upon subscribing the client will be pushed the messages once, further changes are handled
	// // in the PostMessage mutation.
	// r.ChatObservers[id] <- r.ChatMessages

	return nil, nil
}
