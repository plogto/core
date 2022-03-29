package service

import (
	"context"
	"errors"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddPost(ctx context.Context, input model.AddPostInput, postID *string) (*model.Post, error) {
	// authentication
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// check parent post
	var parentPost *model.Post
	if postID != nil {
		parentPost, _ = s.Post.GetPostByID(*postID)

		if parentPost == nil {
			return nil, errors.New("access denied")
		}

		followingUser, _ := s.User.GetUserByID(parentPost.UserID)
		if s.CheckUserAccess(user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}
	}

	// check is empty
	if (input.Attachment == nil || len(input.Attachment) < 1) &&
		(input.Content == nil || len(*input.Content) < 1) {
		return nil, errors.New("need to add attachment or content")
	}

	for _, name := range input.Attachment {
		file, _ := s.File.GetFileByName(name)
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	post := &model.Post{
		ParentID: postID,
		UserID:   user.ID,
		Content:  input.Content,
		Url:      util.RandomString(20),
	}

	s.Post.CreatePost(post)

	// check attachment
	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			s.PostAttachment.CreatePostAttachment(&model.PostAttachment{
				PostID: post.ID,
				Name:   v,
			})
		}
	}

	if len(post.ID) > 0 {
		if post.Content != nil && len(*post.Content) > 0 {
			s.SaveTagsPost(post.ID, *post.Content)
		}
		// notification for reply
		if postID != nil {
			s.CreateNotification(CreateNotificationArgs{
				Name:       constants.NOTIFICATION_REPLY_POST,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "p/" + post.Url + "#" + post.ID,
				PostID:     postID,
				ReplyID:    &post.ID,
			})
		}
	}

	return post, nil
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
				Name:       constants.NOTIFICATION_REPLY_POST,
				SenderID:   user.ID,
				ReceiverID: parentPost.UserID,
				Url:        "p/" + parentPost.Url + "#" + post.ID,
				PostID:     &parentPost.ID,
				ReplyID:    &post.ID,
			})
		}
	}

	s.RemoveNotifications(RemovePostNotificationsArgs{
		ReceiverID: post.UserID,
		PostID:     post.ID,
	})

	return s.Post.DeletePostByID(postID)
}

func (s *Service) GetPostsByParentID(ctx context.Context, parentID string) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	parentPost, _ := s.Post.GetPostByID(parentID)
	followingUser, _ := s.User.GetUserByID(parentPost.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// 	// TODO: add inputPagination
	posts, _ := s.Post.GetPostsByParentIDAndPagination(parentPost.ID, 50, 1)

	return posts, nil
}

func (s *Service) GetPostsByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	followingUser, _ := s.User.GetUserByUsername(username)

	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	var limit int = constants.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	posts, _ := s.Post.GetPostsByUserIDAndPagination(followingUser.ID, nil, limit, page)

	return posts, nil
}

func (s *Service) GetPostsByTagName(ctx context.Context, tagName string, input *model.PaginationInput) (*model.Posts, error) {
	tag, _ := s.Tag.GetTagByName(tagName)

	var limit int = constants.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}

	posts, _ := s.Post.GetPostsByTagIDAndPagination(tag.ID, limit, page)

	return posts, nil
}

func (s *Service) GetPostsCount(ctx context.Context, userID string) (*int, error) {
	count, _ := s.Post.CountPostsByUserID(userID)

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
