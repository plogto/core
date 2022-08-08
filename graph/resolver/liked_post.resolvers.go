package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// User is the resolver for the user field.
func (r *likedPostResolver) User(ctx context.Context, obj *model.LikedPost) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Post is the resolver for the post field.
func (r *likedPostResolver) Post(ctx context.Context, obj *model.LikedPost) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

// LikePost is the resolver for the likePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.LikedPost, error) {
	return r.Service.LikePost(ctx, postID)
}

// GetLikedPostsByPostID is the resolver for the getLikedPostsByPostId field.
func (r *queryResolver) GetLikedPostsByPostID(ctx context.Context, postID string, pageInfoInput *model.PageInfoInput) (*model.LikedPosts, error) {
	return r.Service.GetLikedPostsByPostID(ctx, postID)
}

// LikedPost returns generated.LikedPostResolver implementation.
func (r *Resolver) LikedPost() generated.LikedPostResolver { return &likedPostResolver{r} }

type likedPostResolver struct{ *Resolver }
