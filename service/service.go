package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
)

type Service struct {
	Users               database.Users
	Passwords           database.Passwords
	Posts               database.Posts
	Files               database.Files
	Connections         database.Connections
	Tags                database.Tags
	PostAttachments     database.PostAttachments
	PostTags            database.PostTags
	LikedPosts          database.LikedPosts
	SavedPosts          database.SavedPosts
	OnlineUsers         database.OnlineUsers
	Notifications       database.Notifications
	NotificationTypes   database.NotificationTypes
	OnlineNotifications map[string]chan *model.Notification
	mu                  sync.Mutex
}

func New(service Service) *Service {
	return &Service{
		Users:               service.Users,
		Passwords:           service.Passwords,
		Posts:               service.Posts,
		Files:               service.Files,
		Connections:         service.Connections,
		Tags:                service.Tags,
		PostTags:            service.PostTags,
		PostAttachments:     service.PostAttachments,
		LikedPosts:          service.LikedPosts,
		SavedPosts:          service.SavedPosts,
		OnlineUsers:         service.OnlineUsers,
		Notifications:       service.Notifications,
		NotificationTypes:   service.NotificationTypes,
		OnlineNotifications: map[string]chan *model.Notification{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
