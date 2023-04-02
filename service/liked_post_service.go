package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) LikePost(ctx context.Context, postID uuid.UUID) (*db.LikedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID.String())
	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	likedPost, _ := s.LikedPosts.GetLikedPostByUserIDAndPostID(ctx, user.ID, postID)

	if !validation.IsLikedPostExists(likedPost) {

		likedPost, err := s.LikedPosts.CreateLikedPost(ctx, user.ID, postID)

		if validation.IsLikedPostExists(likedPost) {
			var name = db.NotificationTypeNameLikePost
			if post.ParentID.Valid {
				name = db.NotificationTypeNameLikeReply
			}
			s.CreateNotification(ctx, CreateNotificationArgs{
				Name:       name,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "/p/" + post.Url,
				PostID:     uuid.NullUUID{post.ID, true},
			})
		}

		return likedPost, err

	} else {
		unlikedPost, err := s.LikedPosts.DeleteLikedPostByID(ctx, likedPost.ID)
		s.RemoveNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameLikePost,
			SenderID:   user.ID,
			ReceiverID: post.UserID,
			Url:        "/p/" + post.Url,
			PostID:     uuid.NullUUID{post.ID, true},
		})

		return unlikedPost, err
	}
}

func (s *Service) GetLikedPostsByPostID(ctx context.Context, postID uuid.UUID) (*model.LikedPosts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID.String())

	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) || !validation.IsUserExists(user) {
		return nil, nil
	} else {
		// TODO: add inputPageInfo
		return s.LikedPosts.GetLikedPostsByPostIDAndPageInfo(ctx, postID, 50, "")
	}
}

func (s *Service) GetLikedPostsByUsername(ctx context.Context, username string, input *model.PageInfoInput) (*model.LikedPosts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}

		pageInfo := util.ExtractPageInfo(input)

		return s.LikedPosts.GetLikedPostsByUserIDAndPageInfo(ctx, followingUser.ID, int32(pageInfo.First), pageInfo.After)
	}
}

func (s *Service) IsPostLiked(ctx context.Context, postID uuid.UUID) (*db.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !validation.IsUserExists(user) {
		return nil, nil
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID.String())

	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, nil
	} else {
		isPostLiked, err := s.LikedPosts.GetLikedPostByUserIDAndPostID(ctx, user.ID, postID)
		if len(isPostLiked.ID) < 1 {
			return nil, nil
		}

		return isPostLiked, err
	}
}

func (s *Service) GetLikedPostByID(ctx context.Context, id *string) (*db.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || !validation.IsUserExists(user) {
		return nil, nil
	}

	ID, _ := uuid.Parse(*id)
	likedPost, err := s.LikedPosts.GetLikedPostByID(ctx, ID)
	post, _ := graph.GetPostLoader(ctx).Load(likedPost.PostID.String())

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID.String()); s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, err
	}

	return likedPost, err

}
