package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// SavePost is the resolver for the savePost field.
func (r *mutationResolver) SavePost(ctx context.Context, postID uuid.UUID) (*db.SavedPost, error) {
	return r.Service.SavePost(ctx, postID)
}

// GetSavedPosts is the resolver for the getSavedPosts field.
func (r *queryResolver) GetSavedPosts(ctx context.Context, pageInfo *model.PageInfoInput) (*model.SavedPosts, error) {
	return r.Service.GetSavedPosts(ctx, pageInfo)
}

// User is the resolver for the user field.
func (r *savedPostResolver) User(ctx context.Context, obj *db.SavedPost) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Post is the resolver for the post field.
func (r *savedPostResolver) Post(ctx context.Context, obj *db.SavedPost) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, uuid.NullUUID{obj.PostID, true})
}

// Cursor is the resolver for the cursor field.
func (r *savedPostsEdgeResolver) Cursor(ctx context.Context, obj *model.SavedPostsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *savedPostsEdgeResolver) Node(ctx context.Context, obj *model.SavedPostsEdge) (*db.SavedPost, error) {
	return r.Service.GetSavedPostByID(ctx, obj.Node.ID)
}

// SavedPost returns generated.SavedPostResolver implementation.
func (r *Resolver) SavedPost() generated.SavedPostResolver { return &savedPostResolver{r} }

// SavedPostsEdge returns generated.SavedPostsEdgeResolver implementation.
func (r *Resolver) SavedPostsEdge() generated.SavedPostsEdgeResolver {
	return &savedPostsEdgeResolver{r}
}

type savedPostResolver struct{ *Resolver }
type savedPostsEdgeResolver struct{ *Resolver }
