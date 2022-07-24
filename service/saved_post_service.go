package service

import (
	"context"
	"errors"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) SavePost(ctx context.Context, postID string) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := s.Users.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postSave, _ := s.SavedPosts.GetPostSaveByUserIDAndPostID(user.ID, postID)
	if postSave != nil {
		s.SavedPosts.DeletePostSaveByID(postSave.ID)
	} else {
		postSave := &model.SavedPost{
			UserID: user.ID,
			PostID: postID,
		}

		s.SavedPosts.CreatePostSave(postSave)
	}

	return post, nil
}

func (s *Service) GetSavedPosts(ctx context.Context, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)
	return s.SavedPosts.GetSavedPostsByUserIDAndPagination(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}

func (s *Service) IsPostSaved(ctx context.Context, postID string) (*model.SavedPost, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	if user == nil {
		return nil, nil
	}

	post, _ := s.Posts.GetPostByID(postID)
	followingUser, _ := s.Users.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, nil
	} else {
		return s.SavedPosts.GetPostSaveByUserIDAndPostID(user.ID, postID)
	}
}
