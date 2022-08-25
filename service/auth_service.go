package service

import (
	"context"
	"errors"
	"log"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

func (s *Service) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := s.Users.GetUserByUsernameOrEmail(input.Username)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	password, err := s.Passwords.GetPasswordByUserID(user.ID)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	err = password.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (s *Service) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	if _, err := s.Users.GetUserByUsernameOrEmail(input.Email); err == nil {
		return nil, errors.New("email has already been taken")
	}

	user := &model.User{
		Email:          input.Email,
		Username:       util.RandomString(15),
		InvitationCode: util.RandomString(7),
		FullName:       input.FullName,
	}

	newUser, err := s.Users.CreateUser(user)
	if err != nil {
		log.Printf("error while creating a user: %v", err)
		return nil, err
	}

	password := &model.Password{
		UserID: newUser.ID,
	}

	if err = password.HashPassword(input.Password); err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	if _, err := s.Passwords.AddPassword(password); err != nil {
		log.Printf("error white adding password: %v", err)
		return nil, err
	}

	if input.InvitationCode != nil {
		if inviter, err := s.Users.GetUserByInvitationCode(*input.InvitationCode); err == nil {
			s.InvitedUsers.CreateInvitedUser(&model.InvitedUser{
				InviterID: inviter.ID,
				InviteeID: newUser.ID,
			})
			// TODO: add credits
		}
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
