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

	if user.ID == userID {
		return nil, errors.New("can not follow yourself")
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

func (s *Service) UnfollowUser(ctx context.Context, userID string) (*model.Follower, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not unfollow yourself")
	}

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(userID, user.ID)

	if len(follower.ID) < 1 {
		return nil, errors.New("follower not found")
	}

	deletedFollower, _ := s.Follower.DeleteFollower(follower.ID)

	return deletedFollower, nil
}

func (s *Service) AcceptUser(ctx context.Context, userID string) (*model.Follower, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(user.ID, userID)

	if len(follower.ID) < 1 {
		return nil, errors.New("follower not found")
	}

	if *follower.Status == 1 {
		return follower, nil
	}

	status := 1
	follower.Status = &status

	updatedFollower, _ := s.Follower.UpdateFollower(follower)

	return updatedFollower, nil
}
