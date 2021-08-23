package service

import (
	"context"

	"github.com/favecode/note-core/graph/model"
)

func (s *Service) Search(ctx context.Context, expression string) (*model.Search, error) {
	users, _ := s.SearchUser(ctx, expression)

	return &model.Search{
		User: users,
	}, nil
}
