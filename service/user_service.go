package service

import (
	"context"
	"errors"
	"log"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
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

	if input.Username != nil {
		user.Username = *input.Username
		didUpdate = true
	}

	if input.Avatar != nil {
		user.Avatar = input.Avatar
		if len(*input.Avatar) == 0 {
			user.Avatar = nil
		}
		didUpdate = true
	}

	if input.Background != nil {
		user.Background = input.Background
		if len(*input.Background) == 0 {
			user.Background = nil
		}
		didUpdate = true
	}

	if input.FullName != nil {
		user.FullName = *input.FullName
		didUpdate = true
	}

	if input.Bio != nil {
		user.Bio = input.Bio
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

func (s *Service) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.AuthResponse, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	password, _ := s.Password.GetPasswordByUserID(user.ID)

	if err = password.ComparePassword(input.OldPassword); err != nil {
		return nil, errors.New("old password is not valid")
	}

	if err = password.HashPassword(input.NewPassword); err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	if _, err := s.Password.UpdatePassword(password); err != nil {
		log.Printf("error white updating password: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (s *Service) CheckUsername(ctx context.Context, username string) (*model.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, _ := s.User.GetUserByUsername(username)

	return user, nil
}

func (s *Service) CheckEmail(ctx context.Context, email string) (*model.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, _ := s.User.GetUserByEmail(email)

	return user, nil
}
