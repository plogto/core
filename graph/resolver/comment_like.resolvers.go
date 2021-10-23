package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
)

func (r *commentLikeResolver) User(ctx context.Context, obj *model.CommentLike) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentLikeResolver) Comment(ctx context.Context, obj *model.CommentLike) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnlikeComment(ctx context.Context, commentID string) (*model.CommentLike, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCommentLikesByCommentID(ctx context.Context, commentID string, input *model.PaginationInput) (*model.CommentLikes, error) {
	panic(fmt.Errorf("not implemented"))
}

// CommentLike returns generated.CommentLikeResolver implementation.
func (r *Resolver) CommentLike() generated.CommentLikeResolver { return &commentLikeResolver{r} }

type commentLikeResolver struct{ *Resolver }
