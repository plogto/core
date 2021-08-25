package service

import (
	"context"
	"errors"

	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/middleware"
)

func (s *Service) FollowUser(ctx context.Context, userID string) (*model.Follower, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	followingUser, _ := s.User.GetUserByID(userID)

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(userID, user.ID)

	if len(follower.ID) > 0 {
		return follower, nil
	}

	status := 1
	if followingUser.Private == bool(true) {
		status = 0
	}

	newFollower := &model.Follower{
		FollowerID: user.ID,
		UserID:     userID,
		Status:     &status,
	}

	s.Follower.CreateFollower(newFollower)

	return newFollower, nil
}
