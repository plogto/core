package service

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
	"github.com/samber/lo"
)

func (s *Service) AddPost(ctx context.Context, input model.AddPostInput) (*db.Post, error) {
	// authentication
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// check parent post
	var parentPost *db.Post
	if input.ParentID != nil {
		parentPost, _ = graph.GetPostLoader(ctx).Load(input.ParentID.String())

		if parentPost == nil {
			return nil, errors.New("access denied")
		}

		followingUser, _ := graph.GetUserLoader(ctx).Load(parentPost.UserID.String())
		if !s.CheckUserAccess(ctx, user, followingUser) {
			return nil, errors.New("access denied")
		}
	}

	// check is empty
	if (input.Attachment == nil || len(input.Attachment) < 1) &&
		(input.Content == nil || len(*input.Content) < 1) {
		return nil, errors.New("need to add attachment or content")
	}

	for _, id := range input.Attachment {
		file, _ := s.Files.GetFileByID(ctx, uuid.MustParse(id))
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	content, userIDs := s.FormatPostContent(ctx, *input.Content)

	var parentID uuid.NullUUID
	if input.ParentID != nil {
		parentID = uuid.NullUUID{uuid.MustParse(input.ParentID.String()), true}
	}

	var status = db.PostStatusPublic
	if input.Status != nil {
		status = db.PostStatus(*input.Status)
	}

	post, _ := s.Posts.CreatePost(ctx, db.CreatePostParams{
		ParentID: parentID,
		UserID:   user.ID,
		Content:  sql.NullString{content, true},
		Status:   status,
		Url:      util.RandomString(20),
	})

	if validation.IsPostExists(post) {
		s.CreatePostMentions(ctx, userIDs, post.ID)
	}

	// check attachment
	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			s.PostAttachments.CreatePostAttachment(ctx, post.ID, uuid.MustParse(v))
		}
	}

	if validation.IsPostExists(post) {
		if lo.IsNotEmpty(post.Content) {
			s.SaveTagsPost(ctx, post.ID, post.Content.String)
			s.CreatePostMentionNotifications(ctx, CreatePostMentionNotificationsArgs{
				UserIDs:  userIDs,
				SenderID: user.ID,
				Post:     *post,
			})
		}
		// notification for reply
		if lo.IsNotEmpty(input.ParentID) {
			postID, _ := uuid.Parse(input.ParentID.String())
			s.CreateNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameReplyPost,
				SenderID:   user.ID,
				ReceiverID: post.UserID,
				Url:        "/p/" + post.Url + "#" + post.ID.String(),
				PostID:     uuid.NullUUID{postID, true},
				ReplyID:    uuid.NullUUID{post.ID, true},
			})
		}
	}

	return post, nil
}

func (s *Service) EditPost(ctx context.Context, postID uuid.UUID, input model.EditPostInput) (*db.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	didUpdate := false

	if lo.IsNotEmpty(input.Content) {
		content, userIDs := s.FormatPostContent(ctx, *input.Content)
		if post.Content.String != content {
			oldUserIDs := s.ExtractUserIDsFromPostContent(post.Content.String)

			oldUserIDs = lo.Reject(oldUserIDs, func(oldUser uuid.UUID, _ int) bool {
				_, ok := lo.Find(userIDs, func(user uuid.UUID) bool {
					return oldUser == user
				})

				if ok {
					userIDs = lo.Reject(userIDs, func(userID uuid.UUID, _ int) bool {
						return userID == oldUser
					})
				}

				return ok
			})

			// removed users
			s.DeletePostMentions(ctx, oldUserIDs, postID)
			for _, oldUser := range oldUserIDs {
				s.RemoveNotification(ctx, CreateNotificationArgs{
					Name:       db.NotificationTypeNameMentionInPost,
					SenderID:   user.ID,
					ReceiverID: oldUser,
					Url:        "/p/" + post.Url,
					PostID:     uuid.NullUUID{postID, true},
				})
			}
			// added users
			s.CreatePostMentions(ctx, userIDs, postID)
			s.CreatePostMentionNotifications(ctx, CreatePostMentionNotificationsArgs{
				UserIDs:  userIDs,
				SenderID: user.ID,
				Post:     *post,
			})

			s.PostTags.DeletePostTagsByPostID(ctx, post.ID)
			s.SaveTagsPost(ctx, post.ID, content)

			post.Content = sql.NullString{content, true}
			didUpdate = true
		}
	}

	if input.Status != nil && post.Status != db.PostStatus(input.Status.String()) {
		post.Status = db.PostStatus(input.Status.String())
		didUpdate = true
	}

	if !didUpdate {
		return nil, nil
	}

	return s.Posts.UpdatePost(ctx, post)
}

func (s *Service) DeletePost(ctx context.Context, postID uuid.UUID) (*db.Post, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := graph.GetPostLoader(ctx).Load(postID.String())
	if post == nil || post.UserID != user.ID {
		return nil, errors.New("access denied")
	}

	if validation.IsParentPostExists(post) {
		parentPost, _ := graph.GetPostLoader(ctx).Load(post.ParentID.UUID.String())

		if parentPost != nil && len(parentPost.ID) > 0 {
			// remove notification for reply
			s.RemoveNotification(ctx, CreateNotificationArgs{
				Name:       db.NotificationTypeNameReplyPost,
				SenderID:   user.ID,
				ReceiverID: parentPost.UserID,
				Url:        "/p/" + parentPost.Url + "#" + post.ID.String(),
				PostID:     uuid.NullUUID{parentPost.ID, true},
				ReplyID:    uuid.NullUUID{post.ID, true},
			})
		}
	}

	deletedPost, err := s.Posts.DeletePostByID(ctx, postID)

	if validation.IsPostExists(deletedPost) {
		s.Notifications.RemovePostNotificationsByPostID(ctx,
			postID,
		)

		s.PostMentions.DeletePostMentionsByPostID(ctx, postID)
		s.PostTags.DeletePostTagsByPostID(ctx, postID)
	}

	return deletedPost, nil
}

func (s *Service) GetPostsByParentID(ctx context.Context, parentID uuid.UUID) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	parentPost, _ := graph.GetPostLoader(ctx).Load(parentID.String())
	followingUser, _ := graph.GetUserLoader(ctx).Load(parentPost.UserID.String())
	after := time.Now()

	if !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, nil
	} else {
		// TODO: add inputPageInfo
		if validation.IsUserExists(user) {
			return s.Posts.GetPostsByParentIDAndPageInfo(ctx, uuid.NullUUID{user.ID, true}, parentPost.ID, 50, after)
		} else {
			return s.Posts.GetPostsByParentIDAndPageInfo(ctx, uuid.NullUUID{}, parentPost.ID, 50, after)
		}
	}
}

func (s *Service) GetPostsByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if !s.CheckUserAccess(ctx, user, followingUser) {
			return nil, errors.New("access denied")
		}

		pagination := util.ExtractPageInfo(pageInfo)

		return s.Posts.GetPostsByUserIDAndPageInfo(ctx, followingUser.ID, pagination.First, pagination.After)
	}
}

func (s *Service) GetRepliesByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	followingUser, err := s.Users.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, errors.New("user not found")
	} else {
		if !s.CheckUserAccess(ctx, user, followingUser) {
			return nil, errors.New("access denied")
		}

		pagination := util.ExtractPageInfo(pageInfo)

		return s.Posts.GetPostsWithParentIDByUserIDAndPageInfo(ctx, followingUser.ID, pagination.First, pagination.After)
	}
}

func (s *Service) GetPostsByTagName(ctx context.Context, tagName string, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	tag, err := s.Tags.GetTagByName(ctx, tagName)

	if err != nil {
		return nil, errors.New("tag not found")
	} else {
		pagination := util.ExtractPageInfo(pageInfo)
		return s.Posts.GetPostsByTagIDAndPageInfo(ctx, tag.ID, pagination.First, pagination.After)
	}

}

func (s *Service) GetPostsCount(ctx context.Context, userID uuid.UUID) (int64, error) {
	return s.Posts.CountPostsByUserID(ctx, userID)
}

func (s *Service) GetPostByID(ctx context.Context, id uuid.NullUUID) (*db.Post, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !id.Valid {
		return nil, nil
	}

	post, err := graph.GetPostLoader(ctx).Load(id.UUID.String())

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID.String()); !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, err
	}

	return post, err
}

func (s *Service) GetPostContentByPostID(ctx context.Context, id uuid.NullUUID) (*string, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !id.Valid {
		return nil, nil
	}

	post, err := graph.GetPostLoader(ctx).Load(id.UUID.String())

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID.String()); !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, err
	}

	content := s.ParsePostContent(ctx, post.Content.String)

	return &content, err
}

func (s *Service) GetPostByURL(ctx context.Context, url string) (*db.Post, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, err := s.Posts.GetPostByURL(ctx, url)

	if followingUser, err := graph.GetUserLoader(ctx).Load(post.UserID.String()); !s.CheckUserAccess(ctx, user, followingUser) {
		return nil, err
	}

	return post, err
}

func (s *Service) GetTimelinePosts(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	pagination := util.ExtractPageInfo(pageInfo)

	return s.Posts.GetTimelinePostsByPageInfo(ctx, user.ID, pagination.First, pagination.After)
}

func (s *Service) GetExplorePosts(ctx context.Context, input *model.GetExplorePostsInput, pageInfo *model.PageInfoInput) (*model.Posts, error) {
	pagination := util.ExtractPageInfo(pageInfo)

	if input == nil || !*input.IsAttachment {
		return s.Posts.GetExplorePostsByPageInfo(ctx, pagination.First, pagination.After)
	} else {
		return s.Posts.GetExplorePostsWithAttachmentByPageInfo(ctx, pagination.First, pagination.After)
	}
}

func (s *Service) FormatPostContent(ctx context.Context, content string) (string, []uuid.UUID) {
	var userIDs []uuid.UUID
	r := regexp.MustCompile(constants.MENTION_PATTERN)
	mentions := r.FindAllString(content, -1)
	for _, mention := range mentions {
		username := strings.TrimLeft(mention, "@")
		user, _ := s.Users.GetUserByUsername(ctx, username)
		if validation.IsUserExists(user) {
			content = strings.ReplaceAll(content, mention, util.PrepareKeyPattern(user.ID))
			userIDs = append(userIDs, user.ID)
		}
	}

	return content, userIDs
}

func (s *Service) ExtractUserIDsFromPostContent(content string) []uuid.UUID {
	r := regexp.MustCompile(constants.KEY_PATTERN)
	userIDs := convertor.StringsToUUIDs(r.FindAllString(content, -1))
	for i, userID := range userIDs {
		userIDs[i] = uuid.MustParse(strings.Trim(userID.String(), "$_"))
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
