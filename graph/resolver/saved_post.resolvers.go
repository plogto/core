package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// SavePost is the resolver for the savePost field.
func (r *mutationResolver) SavePost(ctx context.Context, postID string) (*model.Post, error) {
	return r.Service.SavePost(ctx, postID)
}

// GetSavedPosts is the resolver for the getSavedPosts field.
func (r *queryResolver) GetSavedPosts(ctx context.Context, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetSavedPosts(ctx, input)
}

// User is the resolver for the user field.
func (r *savedPostResolver) User(ctx context.Context, obj *model.SavedPost) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Post is the resolver for the post field.
func (r *savedPostResolver) Post(ctx context.Context, obj *model.SavedPost) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

// SavedPost returns generated.SavedPostResolver implementation.
func (r *Resolver) SavedPost() generated.SavedPostResolver { return &savedPostResolver{r} }

type savedPostResolver struct{ *Resolver }
