package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// LikePost is the resolver for the likePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.Post, error) {
	return r.Service.LikePost(ctx, postID)
}

// User is the resolver for the user field.
func (r *postLikeResolver) User(ctx context.Context, obj *model.PostLike) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Post is the resolver for the post field.
func (r *postLikeResolver) Post(ctx context.Context, obj *model.PostLike) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

// GetPostLikesByPostID is the resolver for the getPostLikesByPostId field.
func (r *queryResolver) GetPostLikesByPostID(ctx context.Context, postID string, input *model.PaginationInput) (*model.PostLikes, error) {
	return r.Service.GetPostLikesByPostID(ctx, postID)
}

// PostLike returns generated.PostLikeResolver implementation.
func (r *Resolver) PostLike() generated.PostLikeResolver { return &postLikeResolver{r} }

type postLikeResolver struct{ *Resolver }
