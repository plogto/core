package service

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) LikePost(ctx context.Context, postID pgtype.UUID) (*db.LikedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(postID))
	followingUser, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID))
	if !s.CheckUserAccess(ctx, user, followingUser) {
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
				PostID:     post.ID,
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
			PostID:     post.ID,
		})

		return unlikedPost, err
	}
}

func (s *Service) GetLikedPostsByPostID(ctx context.Context, postID pgtype.UUID) (*model.LikedPosts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(postID))
	followingUser, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID))

	if !s.CheckUserAccess(ctx, user, followingUser) || !validation.IsUserExists(user) {
		return nil, nil
	} else {
		// TODO: add inputPageInfo
		after := time.Now()
		return s.LikedPosts.GetLikedPostsByPostIDAndPageInfo(ctx, postID, 50, after)
	}
}

func (s *Service) GetLikedPostsByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.LikedPosts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if !s.CheckUserAccess(ctx, user, followingUser) {
			return nil, errors.New("access denied")
		}

		pagination := util.ExtractPageInfo(pageInfo)

		if !validation.IsUserExists(user) {
			return nil, nil
		} else {
			return s.LikedPosts.GetLikedPostsByUserIDAndPageInfo(ctx, followingUser.ID, pagination.First, pagination.After)
		}
	}
}

func (s *Service) IsPostLiked(ctx context.Context, postID pgtype.UUID) (*db.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !validation.IsUserExists(user) {
		return nil, nil
	}

	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(postID))
	followingUser, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID))

	if !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, nil
	} else {
		isPostLiked, _ := s.LikedPosts.GetLikedPostByUserIDAndPostID(ctx, user.ID, postID)

		if !validation.IsLikedPostExists(isPostLiked) {
			return nil, nil
		}

		return isPostLiked, nil
	}
}

func (s *Service) GetLikedPostByID(ctx context.Context, id pgtype.UUID) (*db.LikedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !id.Valid || !validation.IsUserExists(user) {
		return nil, nil
	}

	likedPost, err := s.LikedPosts.GetLikedPostByID(ctx, id)
	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(likedPost.PostID))

	if followingUser, err := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID)); !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, err
	}

	return likedPost, err

}
