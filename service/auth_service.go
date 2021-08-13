package service

import (
	"context"
	"errors"
	"log"

	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/util"
)

func (s *Service) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := s.User.GetUserByUsernameOrEmail(input.Username)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	password, err := s.Password.GetPasswordByUserID(user.ID)
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
	_, err := s.User.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email has already been taken")
	}

	user := &model.User{
		Email:    input.Email,
		Username: util.RandomString(15),
		Fullname: input.Fullname,
	}

	userTx, err := s.User.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer userTx.Rollback()

	newUser, err := s.User.CreateUser(userTx, user)

	if err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	if err := userTx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
		return nil, err
	}

	password := &model.Password{
		UserID: newUser.ID,
	}

	err = password.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	passwordTx, err := s.Password.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer passwordTx.Rollback()

	if _, err := s.Password.AddPassword(passwordTx, password); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	if err := passwordTx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
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
