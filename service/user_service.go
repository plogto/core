package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
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
	users, err := s.User.GetUsersByUsernameOrFullnameAndPagination(expression+"%", limit, page)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return users, nil
}

func (s *Service) CheckUserAccess(user *model.User, followingUser *model.User) bool {
	if followingUser.IsPrivate == bool(true) {
		if user != nil {
			connection, _ := s.Connection.GetConnection(followingUser.ID, user.ID)

			if followingUser.ID != user.ID {
				if len(connection.ID) < 1 || *connection.Status < 2 {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}
