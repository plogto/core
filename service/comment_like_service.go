package service

import (
	"context"
	"errors"

	"github.com/plogto/core/config"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) LikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	comment, _ := s.Comment.GetCommentByID(commentID)
	followingUser, _ := s.User.GetUserByID(comment.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	post, _ := s.Post.GetPostByID(comment.PostID)
	if post == nil || post.ID != comment.PostID {
		return nil, errors.New("access denied")
	}

	commentLike := &model.CommentLike{
		UserID:    user.ID,
		CommentID: commentID,
	}

	s.CommentLike.CreateCommentLike(commentLike)

	if len(commentLike.ID) > 0 {
		s.CreateNotification(CreateNotificationArgs{
			Name:       config.NOTIFICATION_LIKE_COMMENT,
			SenderId:   user.ID,
			ReceiverId: comment.UserID,
			Url:        "p/" + post.Url + "#" + comment.ID,
			PostId:     &post.ID,
			CommentId:  &comment.ID,
		})
	}

	return commentLike, nil
}

func (s *Service) UnlikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	comment, _ := s.Comment.GetCommentByID(commentID)
	followingUser, _ := s.User.GetUserByID(comment.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	post, _ := s.Post.GetPostByID(comment.PostID)
	if post == nil || post.ID != comment.PostID {
		return nil, errors.New("access denied")
	}

	commentLike, _ := s.CommentLike.GetCommentLikeByUserIdAndCommentId(user.ID, commentID)

	if commentLike != nil {
		s.RemoveNotification(CreateNotificationArgs{
			Name:       config.NOTIFICATION_LIKE_COMMENT,
			SenderId:   user.ID,
			ReceiverId: comment.UserID,
			Url:        "p/" + post.Url + "#" + comment.ID,
			PostId:     &post.ID,
			CommentId:  &comment.ID,
		})
	} else {
		return nil, errors.New("like not found")
	}

	return s.CommentLike.DeleteCommentLikeByID(commentLike.ID)
}

func (s *Service) GetCommentLikesByCommentId(ctx context.Context, commentID string) (*model.CommentLikes, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	comment, _ := s.Comment.GetCommentByID(commentID)
	followingUser, _ := s.User.GetUserByID(comment.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	commentLikes, _ := s.CommentLike.GetCommentLikesByCommentIdAndPagination(commentID, 10, 1)

	return commentLikes, nil
}

func (s *Service) IsCommentLiked(ctx context.Context, commentID string) (*model.CommentLike, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	comment, _ := s.Comment.GetCommentByID(commentID)
	followingUser, _ := s.User.GetUserByID(comment.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	commentLike, _ := s.CommentLike.GetCommentLikeByUserIdAndCommentId(user.ID, commentID)

	return commentLike, nil
}
