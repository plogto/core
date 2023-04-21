package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) SavePost(ctx context.Context, postID pgtype.UUID) (*db.SavedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(postID))
	followingUser, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID))
	if !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, errors.New("access denied")
	}

	savedPost, _ := s.SavedPosts.GetSavedPostByUserIDAndPostID(ctx, user.ID, postID)

	if !validation.IsSavedPostExists(savedPost) {
		return s.SavedPosts.CreateSavedPost(ctx, user.ID, postID)
	} else {
		return s.SavedPosts.DeleteSavedPostByID(ctx, savedPost.ID)
	}
}

func (s *Service) GetSavedPosts(ctx context.Context, pageInfo *model.PageInfoInput) (*model.SavedPosts, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if !validation.IsUserExists(user) {
		return nil, nil
	}

	pagination := util.ExtractPageInfo(pageInfo)
	return s.SavedPosts.GetSavedPostsByUserIDAndPageInfo(ctx, user.ID, pagination.First, pagination.After)
}

func (s *Service) GetSavedPostByID(ctx context.Context, id pgtype.UUID) (*db.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	savedPost, _ := s.SavedPosts.GetSavedPostByID(ctx, id)
	post, err := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(savedPost.PostID))

	if followingUser, err := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID)); !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, err
	}

	return savedPost, err
}

func (s *Service) IsPostSaved(ctx context.Context, postID pgtype.UUID) (*db.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	post, _ := graph.GetPostLoader(ctx).Load(convertor.UUIDToString(postID))
	followingUser, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(post.UserID))
	if !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, nil
	} else {
		savedPost, err := s.SavedPosts.GetSavedPostByUserIDAndPostID(ctx, user.ID, postID)

		if !validation.IsSavedPostExists(savedPost) {
			return nil, nil
		}

		return savedPost, err
	}
}
