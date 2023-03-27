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

func (s *Service) SavePost(ctx context.Context, postID uuid.UUID) (*db.SavedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID.String())
	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	UserID, _ := uuid.Parse(user.ID)

	savedPost, _ := s.SavedPosts.GetSavedPostByUserIDAndPostID(ctx, UserID, postID)

	if !validation.IsSavedPostExists(savedPost) {
		return s.SavedPosts.CreateSavedPost(ctx, UserID, postID)
	} else {
		return s.SavedPosts.DeleteSavedPostByID(ctx, savedPost.ID)
	}
}

func (s *Service) GetSavedPosts(ctx context.Context, input *model.PageInfoInput) (*model.SavedPosts, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if user == nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)
	return s.SavedPosts.GetSavedPostsByUserIDAndPageInfo(ctx, user.ID, int32(pageInfoInput.First), pageInfoInput.After)
}

func (s *Service) GetSavedPostByID(ctx context.Context, id uuid.UUID) (*db.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	savedPost, _ := s.SavedPosts.GetSavedPostByID(ctx, id)
	post, err := graph.GetPostLoader(ctx).Load(savedPost.PostID.String())

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID.String()); s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, err
	}

	return savedPost, err
}

func (s *Service) IsPostSaved(ctx context.Context, postID uuid.UUID) (*db.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID.String())
	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, nil
	} else {
		UserID, _ := uuid.Parse(user.ID)
		savedPost, err := s.SavedPosts.GetSavedPostByUserIDAndPostID(ctx, UserID, postID)

		if !validation.IsSavedPostExists(savedPost) {
			return nil, nil
		}

		return savedPost, err
	}
}
