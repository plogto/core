package service

import (
	"context"
	"fmt"

	"github.com/favecode/plog-core/database"
	"github.com/favecode/plog-core/graph/model"
)

type Service struct {
	User       database.User
	Password   database.Password
	Post       database.Post
	Connection database.Connection
	Tag        database.Tag
	PostTag    database.PostTag
}

func New(service Service) *Service {
	return &Service{
		User:       service.User,
		Password:   service.Password,
		Post:       service.Post,
		Connection: service.Connection,
		Tag:        service.Tag,
		PostTag:    service.PostTag,
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
