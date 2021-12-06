package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
)

func (r *mutationResolver) SavePost(ctx context.Context, postID string) (*model.PostSave, error) {
	return r.Service.SavePost(ctx, postID)
}

func (r *mutationResolver) UnsavePost(ctx context.Context, postID string) (*model.PostSave, error) {
	return r.Service.UnsavePost(ctx, postID)
}

func (r *postSaveResolver) User(ctx context.Context, obj *model.PostSave) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *postSaveResolver) Post(ctx context.Context, obj *model.PostSave) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, &obj.PostID)
}

func (r *queryResolver) GetSavedPosts(ctx context.Context, input *model.PaginationInput) (*model.PostSaves, error) {
	return r.Service.GetSavedPosts(ctx, input)
}

// PostSave returns generated.PostSaveResolver implementation.
func (r *Resolver) PostSave() generated.PostSaveResolver { return &postSaveResolver{r} }

type postSaveResolver struct{ *Resolver }
