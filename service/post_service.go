package service

import (
	"context"
	"errors"

	"github.com/favecode/poster-core/config"
	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/middleware"
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
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	followingUser, err := s.User.GetUserByUsername(username)

	connection, _ := s.Connection.GetConnection(followingUser.ID, user.ID)

	if followingUser.ID != user.ID {
		if followingUser.Private == bool(true) {
			if len(connection.ID) < 1 || *connection.Status < 2 {
				return nil, errors.New("you need to follow this user")
			}
		}
	}

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

	posts, _ := s.Post.GetPostsByUserIdAndPagination(followingUser.ID, limit, page)

	return posts, nil
}
