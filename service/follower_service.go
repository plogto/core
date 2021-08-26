package service

import (
	"context"
	"errors"

	"github.com/favecode/note-core/config"
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

	if user.ID == userID {
		return nil, errors.New("can not accept yourself")
	}

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(user.ID, userID)

	if len(follower.ID) < 1 {
		return nil, errors.New("follower not found")
	}

	if *follower.Status == 1 {
		return follower, nil
	}

	*follower.Status = 1

	updatedFollower, _ := s.Follower.UpdateFollower(follower)

	return updatedFollower, nil
}

func (s *Service) RejectUser(ctx context.Context, userID string) (*model.Follower, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user.ID == userID {
		return nil, errors.New("can not reject yourself")
	}

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(user.ID, userID)

	if len(follower.ID) < 1 {
		return nil, errors.New("follower not found")
	}

	if *follower.Status == 1 {
		return nil, errors.New("can not reject accepted request")
	}

	deletedFollower, _ := s.Follower.DeleteFollower(follower.ID)

	return deletedFollower, nil
}

func (s *Service) GetUserFollowersByUsername(ctx context.Context, username string, input *model.GetUserFollowersByUserIDInput, resultType string) (*model.Followers, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	followingUser, _ := s.User.GetUserByUsername(username)

	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	follower, _ := s.Follower.GetFollowerByUserIdAndFollowerId(followingUser.ID, user.ID)
	if followingUser.ID != user.ID {
		if followingUser.Private == bool(true) {
			if len(follower.ID) < 1 || *follower.Status == 0 {
				return nil, errors.New("you need to follow this user")
			}
		}
	}

	if resultType == "followers" {
		followers, _ := s.Follower.GetFollowersByUserIdAndPagination(followingUser.ID, limit, page)
		return followers, nil
	} else if resultType == "following" {
		followers, _ := s.Follower.GetFollowingByUserIdAndPagination(followingUser.ID, limit, page)
		return followers, nil
	}

	return nil, nil
}
