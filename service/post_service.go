package service

import (
	"context"
	"errors"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post := &model.Post{
		UserID:  user.ID,
		Content: input.Content,
		Url:     util.RandomString(20),
	}

	s.Post.CreatePost(post)
	s.SaveTagsPost(post.ID, input.Content)

	return post, nil
}

func (s *Service) ReplyPost(ctx context.Context, postID string, input model.AddPostInput) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)

	if post == nil {
		return nil, errors.New("access denied")
	}

	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	reply := &model.Post{
		ParentID: &postID,
		UserID:   user.ID,
		Content:  input.Content,
		Url:      util.RandomString(20),
	}

	s.Post.CreatePost(reply)

	if len(reply.ID) > 0 {
		s.SaveTagsPost(reply.ID, input.Content)

		s.CreateNotification(CreateNotificationArgs{
			Name:       config.NOTIFICATION_REPLY_POST,
			SenderId:   user.ID,
			ReceiverId: post.UserID,
			Url:        "p/" + post.Url + "#" + reply.ID,
			PostId:     &postID,
			ReplyId:    &reply.ID,
		})
	}

	return reply, nil
}

func (s *Service) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(postID)
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	if post.ParentID != nil {
		parentPost, _ := s.Post.GetPostByID(*post.ParentID)

		if parentPost != nil && len(parentPost.ID) > 0 {
			// remove notification for reply
			s.RemoveNotification(CreateNotificationArgs{
				Name:       config.NOTIFICATION_REPLY_POST,
				SenderId:   user.ID,
				ReceiverId: parentPost.UserID,
				Url:        "p/" + parentPost.Url + "#" + post.ID,
				PostId:     &parentPost.ID,
				ReplyId:    &post.ID,
			})
		}
	}

	s.RemoveNotifications(RemovePostNotificationsArgs{
		ReceiverId: post.UserID,
		PostId:     post.ID,
	})

	return s.Post.DeletePostByID(postID)
}

func (s *Service) GetPostsByParentId(ctx context.Context, parentID string) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	parentPost, _ := s.Post.GetPostByID(parentID)
	followingUser, _ := s.User.GetUserByID(parentPost.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// 	// TODO: add inputPagination
	posts, _ := s.Post.GetPostsByParentIdAndPagination(parentPost.ID, 50, 1)

	return posts, nil
}

func (s *Service) GetPostsByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Posts, error) {
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

	posts, _ := s.Post.GetPostsByUserIdAndPagination(followingUser.ID, nil, limit, page)

	return posts, nil
}

func (s *Service) GetPostsByTagName(ctx context.Context, tagName string, input *model.PaginationInput) (*model.Posts, error) {
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

func (s *Service) GetPostByID(ctx context.Context, id *string) (*model.Post, error) {
	if id == nil {
		return nil, nil
	}
	post, _ := s.Post.GetPostByID(*id)

	return post, nil
}

func (s *Service) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	post, _ := s.Post.GetPostByURL(url)

	return post, nil
}
