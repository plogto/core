package service

import (
	"context"
	"fmt"

	"github.com/favecode/note-core/database"
	"github.com/favecode/note-core/graph/model"
)

type Service struct {
	User     database.User
	Password database.Password
	Post     database.Post
}

func New(service Service) *Service {
	return &Service{
		User:     service.User,
		Password: service.Password,
		Post:     service.Post,
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
