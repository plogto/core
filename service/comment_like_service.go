package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
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

	commentLike := &model.CommentLike{
		UserID:    user.ID,
		CommentID: commentID,
	}

	s.CommentLike.CreateCommentLike(commentLike)

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

	commentLike, _ := s.CommentLike.GetCommentLikeByUserIdAndCommentId(user.ID, commentID)

	if len(commentLike.ID) < 1 {
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
