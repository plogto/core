package service

import (
	"context"
	"errors"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) SavePost(ctx context.Context, postID string) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postSave, _ := s.PostSave.GetPostSaveByUserIDAndPostID(user.ID, postID)

	if postSave != nil {
		s.PostSave.DeletePostSaveByID(postSave.ID)
	} else {
		postSave := &model.PostSave{
			UserID: user.ID,
			PostID: postID,
		}

		s.PostSave.CreatePostSave(postSave)
	}

	return post, nil
}

func (s *Service) GetSavedPosts(ctx context.Context, input *model.PaginationInput) (*model.PostSaves, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	postSaves, _ := s.PostSave.GetPostSavesByUserIDAndPagination(user.ID, limit, page)

	return postSaves, nil
}

func (s *Service) IsPostSaved(ctx context.Context, postID string) (*model.PostSave, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postSave, _ := s.PostSave.GetPostSaveByUserIDAndPostID(user.ID, postID)

	return postSave, nil
}
