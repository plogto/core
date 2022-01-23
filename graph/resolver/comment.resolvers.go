package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *commentResolver) Parent(ctx context.Context, obj *model.Comment) (*model.Comment, error) {
	return r.Service.GetCommentByID(ctx, obj.ParentID)
}

func (r *commentResolver) Children(ctx context.Context, obj *model.Comment) (*model.Comments, error) {
	return r.Service.GetChildrenComments(ctx, obj.PostID, obj.ID)
}

func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *commentResolver) Post(ctx context.Context, obj *model.Comment) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

func (r *commentResolver) IsLiked(ctx context.Context, obj *model.Comment) (*model.CommentLike, error) {
	return r.Service.IsCommentLiked(ctx, obj.ID)
}

func (r *mutationResolver) AddComment(ctx context.Context, input model.CommentPostInput) (*model.Comment, error) {
	return r.Service.AddComment(ctx, input)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, commentID string) (*model.Comment, error) {
	return r.Service.DeleteComment(ctx, commentID)
}

func (r *queryResolver) GetComments(ctx context.Context, postID string, input *model.PaginationInput) (*model.Comments, error) {
	return r.Service.GetComments(ctx, postID)
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
