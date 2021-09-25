package service

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/favecode/plog-core/config"
	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
	"github.com/favecode/plog-core/util"
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
	r := regexp.MustCompile("#(\\w|_)+")
	tags := r.FindAllString(input.Content, -1)
	for i, tag := range tags {
		tags[i] = strings.TrimLeft(tag, "#")
	}
	for _, tagName := range util.UniqueSliceElement(tags) {
		tag := &model.Tag{
			Name: tagName,
		}
		s.Tag.CreateTag(tag)

		postTag := &model.PostTag{
			TagID:  tag.ID,
			PostID: post.ID,
		}
		s.PostTag.CreatePostTag(postTag)
	}

	return post, nil
}

func (s *Service) GetUserPostsByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	followingUser, _ := s.User.GetUserByUsername(username)

	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
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

func (s *Service) GetUserPostsByTagName(ctx context.Context, tagName string, input *model.PaginationInput) (*model.Posts, error) {
	tag, _ := s.Tag.GetTagByName(tagName)

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

	posts, _ := s.Post.GetPostsByTagIdAndPagination(tag.ID, limit, page)

	return posts, nil
}

func (s *Service) GetPostsCount(ctx context.Context, userId string) (*int, error) {
	count, _ := s.Post.CountPostsByUserId(userId)

	return count, nil
}

func (s *Service) GetUserPostsByID(ctx context.Context, postId string) (*model.Post, error) {
	post, _ := s.Post.GetPostByID(postId)

	return post, nil
}
