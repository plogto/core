package service

import (
	"context"
	"errors"

	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/validation"
)

func (s *Service) LikePost(ctx context.Context, postID string) (*model.LikedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	likedPost, _ := s.LikedPosts.GetLikedPostByUserIDAndPostID(user.ID, postID)

	if !validation.IsLikedPostExists(likedPost) {
		likedPost, err := s.LikedPosts.CreateLikedPost(&model.LikedPost{
			UserID: user.ID,
			PostID: postID,
		})

		if validation.IsLikedPostExists(likedPost) {
			var name = model.NotificationTypeNameLikePost
			if post.ParentID != nil {
				name = model.NotificationTypeNameLikeReply
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
			Name:       model.NotificationTypeNameLikePost,
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

	post, _ := graph.GetPostLoader(ctx).Load(postID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID)

	if s.CheckUserAccess(user, followingUser) == bool(false) || !validation.IsUserExists(user) {
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

	post, _ := graph.GetPostLoader(ctx).Load(postID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID)
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

func (s *Service) GetLikedPostByID(ctx context.Context, id *string) (*model.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || !validation.IsUserExists(user) {
		return nil, nil
	}

	likedPost, err := s.LikedPosts.GetLikedPostByID(*id)
	post, _ := graph.GetPostLoader(ctx).Load(likedPost.PostID)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID); s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, err
	}

	return likedPost, err

}
