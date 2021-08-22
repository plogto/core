package service

import (
	"context"
	"errors"

	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/middleware"
)

func (s *Service) GetUserInfo(ctx context.Context) (*model.User, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.User.GetUserByID(id)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := s.User.GetUserByUsername(username)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *Service) SearchUser(ctx context.Context, expression string) (*model.Users, error) {
	limit := 10
	page := 1
	users, err := s.User.GetUserByUsernameOrFullnameAndPagination(expression+"%", limit, page)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return users, nil
}
