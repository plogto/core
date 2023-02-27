package service

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
	"github.com/samber/lo"
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
		parentPost, _ = graph.GetPostLoader(ctx).Load(*input.ParentID)

		if parentPost == nil {
			return nil, errors.New("access denied")
		}

		followingUser, _ := graph.GetUserLoader(ctx).Load(parentPost.UserID)
		if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}
	}

	// check is empty
	if (input.Attachment == nil || len(input.Attachment) < 1) &&
		(input.Content == nil || len(*input.Content) < 1) {
		return nil, errors.New("need to add attachment or content")
	}

	for _, id := range input.Attachment {
		ID, _ := uuid.Parse(id)
		file, _ := s.Files.GetFileByID(ctx, ID)
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	content, userIDs := s.FormatPostContent(*input.Content)

	post := &model.Post{
		ParentID: input.ParentID,
		UserID:   user.ID,
		Content:  &content,
		Status:   input.Status,
		Url:      util.RandomString(20),
	}

	s.Posts.CreatePost(post)

	if validation.IsPostExists(post) {
		s.CreatePostMentions(userIDs, post.ID)
	}

	// check attachment
	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			PostID, _ := uuid.Parse(post.ID)
			V, _ := uuid.Parse(v)
			s.PostAttachments.CreatePostAttachment(ctx, db.CreatePostAttachmentParams{
				PostID: PostID,
				FileID: V,
			})
		}
	}

	if validation.IsPostExists(post) {
		if lo.IsNotEmpty(post.Content) {
			s.SaveTagsPost(post.ID, *post.Content)
			s.CreatePostMentionNotifications(ctx, CreatePostMentionNotificationsArgs{
				UserIDs:  userIDs,
				SenderID: user.ID,
				Post:     *post,
			})
		}
		// notification for reply
		if lo.IsNotEmpty(input.ParentID) {
			// FIXME
			senderID, _ := uuid.Parse(user.ID)
			receiverID, _ := uuid.Parse(post.UserID)
			postID, _ := uuid.Parse(*input.ParentID)
			replyID, _ := uuid.Parse(post.ID)
			s.CreateNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameReplyPost,
				SenderID:   senderID,
				ReceiverID: receiverID,
				Url:        "/p/" + post.Url + "#" + post.ID,
				PostID:     uuid.NullUUID{postID, true},
				ReplyID:    uuid.NullUUID{replyID, true},
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

	post, _ := graph.GetPostLoader(ctx).Load(postID)
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	didUpdate := false

	if lo.IsNotEmpty(input.Content) {
		content, userIDs := s.FormatPostContent(*input.Content)
		if post.Content != &content {
			oldUserIDs := s.ExtractUserIDsFromPostContent(*post.Content)

			oldUserIDs = lo.Reject(oldUserIDs, func(oldUser string, _ int) bool {
				_, ok := lo.Find(userIDs, func(user string) bool {
					return oldUser == user
				})

				if ok {
					userIDs = lo.Reject(userIDs, func(userID string, _ int) bool {
						return userID == oldUser
					})
				}

				return ok
			})

			// removed users
			s.DeletePostMentions(oldUserIDs, postID)
			for _, oldUser := range oldUserIDs {
				// FIXME
				senderID, _ := uuid.Parse(user.ID)
				receiverID, _ := uuid.Parse(oldUser)
				postID, _ := uuid.Parse(postID)
				s.RemoveNotification(ctx, CreateNotificationArgs{
					Name:       db.NotificationTypeNameMentionInPost,
					SenderID:   senderID,
					ReceiverID: receiverID,
					Url:        "/p/" + post.Url,
					PostID:     uuid.NullUUID{postID, true},
				})
			}
			// added users
			s.CreatePostMentions(userIDs, postID)
			s.CreatePostMentionNotifications(ctx, CreatePostMentionNotificationsArgs{
				UserIDs:  userIDs,
				SenderID: user.ID,
				Post:     *post,
			})

			s.PostTags.DeletePostTagsByPostID(post.ID)
			s.SaveTagsPost(post.ID, content)

			post.Content = &content
			didUpdate = true
		}
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

	post, _ := graph.GetPostLoader(ctx).Load(postID)
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	if post.ParentID != nil {
		parentPost, _ := graph.GetPostLoader(ctx).Load(*post.ParentID)

		if parentPost != nil && len(parentPost.ID) > 0 {
			// remove notification for reply
			// FIXME
			senderID, _ := uuid.Parse(user.ID)
			receiverID, _ := uuid.Parse(parentPost.UserID)
			postID, _ := uuid.Parse(parentPost.ID)
			replyID, _ := uuid.Parse(post.ID)
			s.RemoveNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameReplyPost,
				SenderID:   senderID,
				ReceiverID: receiverID,
				Url:        "/p/" + parentPost.Url + "#" + post.ID,
				PostID:     uuid.NullUUID{postID, true},
				ReplyID:    uuid.NullUUID{replyID, true},
			})
		}
	}

	deletedAt := time.Now()
	deletedPost, err := s.Posts.DeletePostByID(postID)

	if err == nil {
		s.Notifications.RemovePostNotificationsByPostID(ctx,
			postID,
		)

		s.PostMentions.DeletePostMentionsByPostID(&model.PostMention{
			PostID:    postID,
			DeletedAt: &deletedAt,
		})
	}

	return deletedPost, nil
}

func (s *Service) GetPostsByParentID(ctx context.Context, parentID string) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	parentPost, _ := graph.GetPostLoader(ctx).Load(parentID)
	followingUser, _ := graph.GetUserLoader(ctx).Load(parentPost.UserID)

	if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, nil
	} else {
		// TODO: add inputPageInfo
		if validation.IsUserExists(user) {
			return s.Posts.GetPostsByParentIDAndPageInfo(&user.ID, parentPost.ID, 50, "")
		} else {
			return s.Posts.GetPostsByParentIDAndPageInfo(nil, parentPost.ID, 50, "")
		}
	}
}

func (s *Service) GetPostsByUsername(ctx context.Context, username string, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}

		pageInfoInput := util.ExtractPageInfo(input)

		return s.Posts.GetPostsByUserIDAndPageInfo(followingUser.ID, nil, pageInfoInput.First, pageInfoInput.After)
	}
}

func (s *Service) GetRepliesByUsername(ctx context.Context, username string, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
			return nil, errors.New("access denied")
		}

		pageInfoInput := util.ExtractPageInfo(input)

		return s.Posts.GetPostsWithParentIDByUserIDAndPageInfo(followingUser.ID, pageInfoInput.First, pageInfoInput.After)
	}
}

func (s *Service) GetPostsByTagName(ctx context.Context, tagName string, input *model.PageInfoInput) (*model.Posts, error) {
	tag, err := s.Tags.GetTagByName(tagName)

	if err != nil {
		return nil, errors.New("tag not found")
	} else {
		pageInfoInput := util.ExtractPageInfo(input)
		return s.Posts.GetPostsByTagIDAndPageInfo(tag.ID, pageInfoInput.First, pageInfoInput.After)
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

	post, err := graph.GetPostLoader(ctx).Load(*id)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID); s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, err
	}

	return post, err
}

func (s *Service) GetPostContentByPostID(ctx context.Context, id *string) (*string, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil {
		return nil, nil
	}

	post, err := graph.GetPostLoader(ctx).Load(*id)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID); s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, err
	}

	content := s.ParsePostContent(ctx, *post.Content)

	return &content, err
}

func (s *Service) GetPostByURL(ctx context.Context, url string) (*model.Post, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, err := s.Posts.GetPostByURL(url)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID); s.CheckUserAccess(ctx, user, followingUser) == bool(false) {
		return nil, err
	}

	return post, err
}

func (s *Service) GetTimelinePosts(ctx context.Context, input *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	pageInfoInput := util.ExtractPageInfo(input)

	// FIXME
	return s.Posts.GetTimelinePostsByPageInfo(user.ID, pageInfoInput.First, pageInfoInput.After)
}

func (s *Service) GetExplorePosts(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.Posts, error) {
	pageInfo := util.ExtractPageInfo(pageInfoInput)

	// FIXME
	return s.Posts.GetExplorePostsByPageInfo(pageInfo.First, pageInfo.After)
}

func (s *Service) FormatPostContent(content string) (string, []string) {
	var userIDs []string
	r := regexp.MustCompile(constants.MENTION_PATTERN)
	mentions := r.FindAllString(content, -1)
	for _, mention := range mentions {
		username := strings.TrimLeft(mention, "@")
		user, _ := s.Users.GetUserByUsername(username)
		if validation.IsUserExists(user) {
			content = strings.ReplaceAll(content, mention, util.PrepareKeyPattern(user.ID))
			userIDs = append(userIDs, user.ID)
		}
	}

	return content, userIDs
}

func (s *Service) ExtractUserIDsFromPostContent(content string) []string {
	r := regexp.MustCompile(constants.KEY_PATTERN)
	userIDs := r.FindAllString(content, -1)
	for i, userID := range userIDs {
		userIDs[i] = strings.Trim(userID, "$_")
	}

	return userIDs
}

func (s *Service) ParsePostContent(ctx context.Context, content string) string {
	r := regexp.MustCompile(constants.KEY_PATTERN)
	userIDKeys := r.FindAllString(content, -1)
	for _, userIDKey := range userIDKeys {
		user, _ := graph.GetUserLoader(ctx).Load(strings.Trim(userIDKey, "$_"))
		if validation.IsUserExists(user) {
			content = strings.ReplaceAll(content, userIDKey, "@"+user.Username)
		}
	}

	return content
}
