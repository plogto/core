package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) LikePost(ctx context.Context, postID string) (*model.LikedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := s.Users.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	likedPost, _ := s.LikedPosts.GetLikedPostByUserIDAndPostID(user.ID, postID)

	if util.IsEmpty(likedPost.ID) {
		likedPost, err := s.LikedPosts.CreateLikedPost(&model.LikedPost{
			UserID: user.ID,
			PostID: postID,
		})

		fmt.Println(likedPost)

		if !util.IsEmpty(likedPost.ID) {
			var name string = constants.NOTIFICATION_LIKE_POST
			if post.ParentID != nil {
				name = constants.NOTIFICATION_LIKE_REPLY
			}

			s.CreateNotification(CreateNotificationArgs{
				Name:       name,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "/p/" + post.Url,
				PostID:     &post.ID,
			})
		}

		return likedPost, err

	} else {
		unlikedPost, err := s.LikedPosts.DeleteLikedPostByID(likedPost.ID)

		s.RemoveNotification(CreateNotificationArgs{
			Name:       constants.NOTIFICATION_LIKE_POST,
			SenderID:   user.ID,
			ReceiverID: post.UserID,
			Url:        "/p/" + post.Url,
			PostID:     &post.ID,
		})

		return unlikedPost, err
	}
}

func (s *Service) GetLikedPostsByPostID(ctx context.Context, postID string) (*model.LikedPosts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := s.Users.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, nil
	} else {
		// TODO: add inputPageInfo
		return s.LikedPosts.GetLikedPostsByPostIDAndPageInfo(postID, 50, "")
	}
}

func (s *Service) IsPostLiked(ctx context.Context, postID string) (*model.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := s.Users.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, nil
	} else {
		isPostLiked, err := s.LikedPosts.GetLikedPostByUserIDAndPostID(user.ID, postID)
		if len(isPostLiked.ID) < 1 {
			return nil, nil
		}

		return isPostLiked, err
	}
}
