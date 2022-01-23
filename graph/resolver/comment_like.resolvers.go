package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *commentLikeResolver) User(ctx context.Context, obj *model.CommentLike) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *commentLikeResolver) Comment(ctx context.Context, obj *model.CommentLike) (*model.Comment, error) {
	return r.Service.GetCommentByID(ctx, &obj.CommentID)
}

func (r *mutationResolver) LikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	return r.Service.LikeComment(ctx, commentID)
}

func (r *mutationResolver) UnlikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	return r.Service.UnlikeComment(ctx, commentID)
}

func (r *queryResolver) GetCommentLikesByCommentID(ctx context.Context, commentID string, input *model.PaginationInput) (*model.CommentLikes, error) {
	return r.Service.GetCommentLikesByCommentId(ctx, commentID)
}

// CommentLike returns generated.CommentLikeResolver implementation.
func (r *Resolver) CommentLike() generated.CommentLikeResolver { return &commentLikeResolver{r} }

type commentLikeResolver struct{ *Resolver }
