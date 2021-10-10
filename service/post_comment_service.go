package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/middleware"
)

func (s *Service) AddPostComment(ctx context.Context, input model.CommentPostInput) (*model.PostComment, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	post, _ := s.Post.GetPostByID(input.PostID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	postComment := &model.PostComment{
		UserID:   user.ID,
		PostID:   input.PostID,
		Content:  input.Content,
		ParentID: input.ParentID,
	}

	s.PostComment.CreatePostComment(postComment)

	return postComment, nil
}

func (s *Service) GetChildrenPostComments(ctx context.Context, postID string, parentId string) (*model.PostComments, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	// TODO: add inputPagination
	postComments, _ := s.PostComment.GetPostCommentsByParentIdAndPagination(parentId, 10, 1)

	return postComments, nil
}

func (s *Service) GetPostComments(ctx context.Context, postID string) (*model.PostComments, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	post, _ := s.Post.GetPostByID(postID)
	followingUser, _ := s.User.GetUserByID(post.UserID)
	if s.CheckUserAccess(user, followingUser) == bool(false) {
		return nil, errors.New("access denied")
	}

	fmt.Println("====>>>>", postID)

	// TODO: add inputPagination
	postComments, _ := s.PostComment.GetPostCommentsByPostIdAndPagination(postID, 10, 1)

	return postComments, nil
}

func (s *Service) GetPostCommentByID(ctx context.Context, id *string) (*model.PostComment, error) {
	if id == nil {
		return nil, nil
	}

	postComment, _ := s.PostComment.GetPostCommentByID(*id)

	return postComment, nil
}
