package service

import (
	"context"
	"errors"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) LikePost(ctx context.Context, postID string) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postLike, _ := s.PostLike.GetPostLikeByUserIDAndPostID(user.ID, postID)

	if postLike != nil {
		s.PostLike.DeletePostLikeByID(postLike.ID)
		s.RemoveNotification(CreateNotificationArgs{
			Name:       constants.NOTIFICATION_LIKE_POST,
			SenderID:   user.ID,
			ReceiverID: post.UserID,
			Url:        "p/" + post.Url,
			PostID:     &post.ID,
		})
	} else {
		postLike := &model.PostLike{
			UserID: user.ID,
			PostID: postID,
		}

		s.PostLike.CreatePostLike(postLike)

		if len(postLike.ID) > 0 {
			var name string = constants.NOTIFICATION_LIKE_POST

			if post.ParentID != nil {
				name = constants.NOTIFICATION_LIKE_REPLY
			}

			s.CreateNotification(CreateNotificationArgs{
				Name:       name,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "p/" + post.Url,
				PostID:     &post.ID,
			})
		}
	}

	return post, nil
}

func (s *Service) GetPostLikesByPostID(ctx context.Context, postID string) (*model.PostLikes, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	postLikes, _ := s.PostLike.GetPostLikesByPostIDAndPagination(postID, 10, 1)

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

	postLike, _ := s.PostLike.GetPostLikeByUserIDAndPostID(user.ID, postID)

	return postLike, nil
}
