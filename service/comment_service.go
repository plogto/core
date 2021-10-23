package service

import (
	"context"
	"errors"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) AddComment(ctx context.Context, input model.CommentPostInput) (*model.Comment, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len(input.Content) < 1 {
		return nil, errors.New("content is empty")
	}

	post, _ := s.Post.GetPostByID(input.PostID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	comment := &model.Comment{
		UserID:   user.ID,
		PostID:   input.PostID,
		Content:  input.Content,
		ParentID: input.ParentID,
	}

	s.Comment.CreateComment(comment)

	return comment, nil
}

func (s *Service) GetChildrenComments(ctx context.Context, postID string, parentId string) (*model.Comments, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	comments, _ := s.Comment.GetCommentsByParentIdAndPagination(parentId, 10, 1)

	return comments, nil
}

func (s *Service) GetComments(ctx context.Context, postID string) (*model.Comments, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	comments, _ := s.Comment.GetCommentsByPostIdAndPagination(postID, 10, 1)

	return comments, nil
}

func (s *Service) GetCommentByID(ctx context.Context, id *string) (*model.Comment, error) {
	if id == nil {
		return nil, nil
	}

	comment, _ := s.Comment.GetCommentByID(*id)

	return comment, nil
}
