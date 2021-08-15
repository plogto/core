package service

import (
	"context"
	"errors"
	"fmt"

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

	fmt.Println("====>>>>")
	fmt.Println(user)
	fmt.Println("====>>>>")

	return user, nil
}
