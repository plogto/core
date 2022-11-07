package service

import (
	"context"
	"errors"

	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) SavePost(ctx context.Context, postID string) (*model.SavedPost, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	savedPost, _ := s.SavedPosts.GetSavedPostByUserIDAndPostID(user.ID, postID)

	if !validation.IsSavedPostExists(savedPost) {
		savedPost := &model.SavedPost{
			UserID: user.ID,
			PostID: postID,
		}
		return s.SavedPosts.CreateSavedPost(savedPost)
	} else {
		return s.SavedPosts.DeleteSavedPostByID(savedPost.ID)
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
	return s.SavedPosts.GetSavedPostsByUserIDAndPageInfo(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}

func (s *Service) GetSavedPostByID(ctx context.Context, id string) (*model.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	savedPost, _ := s.SavedPosts.GetSavedPostByID(id)
	post, err := s.Posts.GetPostByID(savedPost.PostID)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID); s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, err
	}

	return savedPost, err
}

func (s *Service) IsPostSaved(ctx context.Context, postID string) (*model.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, nil
	} else {
		savedPost, err := s.SavedPosts.GetSavedPostByUserIDAndPostID(user.ID, postID)

		if !validation.IsSavedPostExists(savedPost) {
			return nil, nil
		}

		return savedPost, err
	}
}
