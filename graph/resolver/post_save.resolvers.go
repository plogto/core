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

// User is the resolver for the user field.
func (r *postSaveResolver) User(ctx context.Context, obj *model.PostSave) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Post is the resolver for the post field.
func (r *postSaveResolver) Post(ctx context.Context, obj *model.PostSave) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

// GetSavedPosts is the resolver for the getSavedPosts field.
func (r *queryResolver) GetSavedPosts(ctx context.Context, input *model.PaginationInput) (*model.Posts, error) {
	return r.Service.GetSavedPosts(ctx, input)
}

// PostSave returns generated.PostSaveResolver implementation.
func (r *Resolver) PostSave() generated.PostSaveResolver { return &postSaveResolver{r} }

type postSaveResolver struct{ *Resolver }
