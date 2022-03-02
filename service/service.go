package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/model"
)

type Service struct {
	User             database.User
	Password         database.Password
	Post             database.Post
	File             database.File
	Connection       database.Connection
	Tag              database.Tag
	PostTag          database.PostTag
	PostLike         database.PostLike
	PostSave         database.PostSave
	OnlineUser       database.OnlineUser
	Notification     database.Notification
	NotificationType database.NotificationType
	Notifications    map[string]chan *model.Notification
	mu               sync.Mutex
}

func New(service Service) *Service {
	return &Service{
		User:             service.User,
		Password:         service.Password,
		Post:             service.Post,
		File:             service.File,
		Connection:       service.Connection,
		Tag:              service.Tag,
		PostTag:          service.PostTag,
		PostLike:         service.PostLike,
		PostSave:         service.PostSave,
		OnlineUser:       service.OnlineUser,
		Notification:     service.Notification,
		NotificationType: service.NotificationType,
		Notifications:    map[string]chan *model.Notification{},
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
