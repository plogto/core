package service

import (
	"context"
	"errors"

	"github.com/favecode/note-core/config"
	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/middleware"
)

func (s *Service) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post := &model.Post{
		UserID:  user.ID,
		Content: input.Content,
		Status:  input.Status,
	}

	s.Post.CreatePost(post)

	return post, nil
}

func (s *Service) GetUserPostsByUsername(ctx context.Context, username string, input *model.GetUserPostsByUsernameInput) (*model.Posts, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, err := s.User.GetUserByUsername(username)

	if err != nil {
		return nil, errors.New("user not found")
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

	posts, _ := s.Post.GetPostsByUserIdAndPagination(user.ID, limit, page)

	return posts, nil
}
