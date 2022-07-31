package service

import (
	"context"
	"errors"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	// authentication
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// check parent post
	var parentPost *model.Post
	if input.ParentID != nil {
		parentPost, _ = s.Posts.GetPostByID(*input.ParentID)

		if parentPost == nil {
			return nil, errors.New("access denied")
		}

		followingUser, _ := s.Users.GetUserByID(parentPost.UserID)
		if s.CheckUserAccess(user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}
	}

	// check is empty
	if (input.Attachment == nil || len(input.Attachment) < 1) &&
		(input.Content == nil || len(*input.Content) < 1) {
		return nil, errors.New("need to add attachment or content")
	}

	for _, id := range input.Attachment {
		file, _ := s.Files.GetFileByID(id)
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	post := &model.Post{
		ParentID: input.ParentID,
		UserID:   user.ID,
		Content:  input.Content,
		Status:   input.Status,
		Url:      util.RandomString(20),
	}
	s.Posts.CreatePost(post)

	// check attachment
	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			s.PostAttachments.CreatePostAttachment(&model.PostAttachment{
				PostID: post.ID,
				FileID: v,
			})
		}
	}

	if len(post.ID) > 0 {
		if post.Content != nil && len(*post.Content) > 0 {
			s.SaveTagsPost(post.ID, *post.Content)
		}
		// notification for reply
		if input.ParentID != nil {
			s.CreateNotification(CreateNotificationArgs{
				Name:       constants.NOTIFICATION_REPLY_POST,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "/p/" + post.Url + "#" + post.ID,
				PostID:     input.ParentID,
				ReplyID:    &post.ID,
			})
		}
	}

	return post, nil
}

func (s *Service) EditPost(ctx context.Context, postID string, input model.EditPostInput) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Posts.GetPostByID(postID)
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	didUpdate := false

	if input.Content != nil && post.Content != input.Content {
		post.Content = input.Content
		didUpdate = true
	}

	if input.Status != nil && post.Status != input.Status {
		post.Status = input.Status
		didUpdate = true
	}

	if didUpdate == bool(false) {
		return nil, nil
	}

	return s.Posts.UpdatePost(post)
}

func (s *Service) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Posts.GetPostByID(postID)
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	if post.ParentID != nil {
		parentPost, _ := s.Posts.GetPostByID(*post.ParentID)

		if parentPost != nil && len(parentPost.ID) > 0 {
			// remove notification for reply
			s.RemoveNotification(CreateNotificationArgs{
				Name:       constants.NOTIFICATION_REPLY_POST,
				SenderID:   user.ID,
				ReceiverID: parentPost.UserID,
				Url:        "/p/" + parentPost.Url + "#" + post.ID,
				PostID:     &parentPost.ID,
				ReplyID:    &post.ID,
			})
		}
	}

	s.RemoveNotifications(RemovePostNotificationsArgs{
		ReceiverID: post.UserID,
		PostID:     post.ID,
	})

	return s.Posts.DeletePostByID(postID)
}

func (s *Service) GetPostsByParentID(ctx context.Context, parentID string) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	parentPost, _ := s.Posts.GetPostByID(parentID)
	followingUser, _ := s.Users.GetUserByID(parentPost.UserID)

	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, nil
	} else {
		// TODO: add inputPagination
		return s.Posts.GetPostsByParentIDAndPagination(parentPost.ID, 50, "")
	}
}

func (s *Service) GetPostsByUsername(ctx context.Context, username string, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if s.CheckUserAccess(user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}

		pageInfoInput := util.ExtractPageInfo(input)

		return s.Posts.GetPostsByUserIDAndPagination(followingUser.ID, nil, *pageInfoInput.First, *pageInfoInput.After)
	}
}

func (s *Service) GetPostsByTagName(ctx context.Context, tagName string, input *model.PageInfoInput) (*model.Posts, error) {
	tag, err := s.Tags.GetTagByName(tagName)

	if err != nil {
		return nil, errors.New("tag not found")
	} else {
		pageInfoInput := util.ExtractPageInfo(input)
		return s.Posts.GetPostsByTagIDAndPagination(tag.ID, *pageInfoInput.First, *pageInfoInput.After)
	}

}

func (s *Service) GetPostsCount(ctx context.Context, userID string) (int, error) {
	return s.Posts.CountPostsByUserID(userID)
}

func (s *Service) GetPostByID(ctx context.Context, id *string) (*model.Post, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil {
		return nil, nil
	}

	post, err := s.Posts.GetPostByID(*id)

	if followingUser, err := s.Users.GetUserByID(post.UserID); s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, err
	}

	return post, err
}

func (s *Service) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	return s.Posts.GetPostByURL(url)
}

func (s *Service) GetTimelinePosts(ctx context.Context, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	pageInfoInput := util.ExtractPageInfo(input)

	return s.Posts.GetTimelinePostsByPagination(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}
