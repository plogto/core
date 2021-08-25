package service

import (
	"context"
	"errors"

	"github.com/favecode/note-core/graph/model"
)

func (s *Service) Search(ctx context.Context, expression string) (*model.Search, error) {
	if len(expression) < 1 {
		return nil, errors.New("expression is not valid")
	}
	users, _ := s.SearchUser(ctx, expression)

	return &model.Search{
		User: users,
	}, nil
}
