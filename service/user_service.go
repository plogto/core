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
	user, _ := s.User.GetUserByID(id)

	return user, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, _ := s.User.GetUserByUsername(username)

	return user, nil
}

func (s *Service) SearchUser(ctx context.Context, expression string) (*model.Users, error) {
	limit := 10
	page := 1
	users, _ := s.User.GetUsersByUsernameOrFullNameAndPagination(expression+"%", limit, page)

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

func (s *Service) EditUser(ctx context.Context, input model.EditUserInput) (*model.User, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	didUpdate := false

	if input.FullName != nil {
		user.FullName = input.FullName
		didUpdate = true
	}

	if input.Email != nil {
		user.Email = *input.Email
		didUpdate = true
	}

	if input.IsPrivate != nil {
		user.IsPrivate = *input.IsPrivate
		didUpdate = true
	}

	if didUpdate == bool(false) {
		return nil, nil
	}

	updatedUser, _ := s.User.UpdateUser(user)

	return updatedUser, nil
}
