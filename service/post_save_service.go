package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/config"
	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) SavePost(ctx context.Context, postID string) (*model.PostSave, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postSave := &model.PostSave{
		UserID: user.ID,
		PostID: postID,
	}

	s.PostSave.CreatePostSave(postSave)

	return postSave, nil
}

func (s *Service) UnsavePost(ctx context.Context, postID string) (*model.PostSave, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postSave, _ := s.PostSave.GetPostSaveByUserIdAndPostId(user.ID, postID)

	if len(postSave.ID) < 1 {
		return nil, errors.New("save not found")
	}

	return s.PostSave.DeletePostSaveByID(postSave.ID)
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

	postSaves, _ := s.PostSave.GetPostSavesByUserIdAndPagination(user.ID, limit, page)

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

	postSave, _ := s.PostSave.GetPostSaveByUserIdAndPostId(user.ID, postID)

	return postSave, nil
}