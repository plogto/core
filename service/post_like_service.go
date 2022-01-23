package service

import (
	"context"
	"errors"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) LikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postLike := &model.PostLike{
		UserID: user.ID,
		PostID: postID,
	}

	s.PostLike.CreatePostLike(postLike)

	if len(postLike.ID) > 0 {
		s.CreateNotification(CreateNotificationArgs{
			Name:       config.NOTIFICATION_LIKE_POST,
			SenderId:   user.ID,
			ReceiverId: post.UserID,
			Url:        "p/" + post.Url,
			PostId:     &post.ID,
		})
	}

	return postLike, nil
}

func (s *Service) UnlikePost(ctx context.Context, postID string) (*model.PostLike, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postLike, _ := s.PostLike.GetPostLikeByUserIdAndPostId(user.ID, postID)

	if postLike != nil {
		s.RemoveNotification(CreateNotificationArgs{
			Name:       config.NOTIFICATION_LIKE_POST,
			SenderId:   user.ID,
			ReceiverId: post.UserID,
			Url:        "p/" + post.Url,
			PostId:     &post.ID,
		})
	} else {
		return nil, errors.New("like not found")
	}

	return s.PostLike.DeletePostLikeByID(postLike.ID)
}

func (s *Service) GetPostLikesByPostId(ctx context.Context, postID string) (*model.PostLikes, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	postLikes, _ := s.PostLike.GetPostLikesByPostIdAndPagination(postID, 10, 1)

	return postLikes, nil
}

func (s *Service) IsPostLiked(ctx context.Context, postID string) (*model.PostLike, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postLike, _ := s.PostLike.GetPostLikeByUserIdAndPostId(user.ID, postID)

	return postLike, nil
}
