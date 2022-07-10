package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// Following is the resolver for the following field.
func (r *connectionResolver) Following(ctx context.Context, obj *model.Connection) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowingID)
}

// Follower is the resolver for the follower field.
func (r *connectionResolver) Follower(ctx context.Context, obj *model.Connection) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowerID)
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.FollowUser(ctx, userID)
}

// UnfollowUser is the resolver for the unfollowUser field.
func (r *mutationResolver) UnfollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.UnfollowUser(ctx, userID)
}

// AcceptUser is the resolver for the acceptUser field.
func (r *mutationResolver) AcceptUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.AcceptUser(ctx, userID)
}

// RejectUser is the resolver for the rejectUser field.
func (r *mutationResolver) RejectUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.RejectUser(ctx, userID)
}

// GetFollowersByUsername is the resolver for the getFollowersByUsername field.
func (r *queryResolver) GetFollowersByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, input, "followers")
}

// GetFollowingByUsername is the resolver for the getFollowingByUsername field.
func (r *queryResolver) GetFollowingByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, input, "following")
}

// GetFollowRequests is the resolver for the getFollowRequests field.
func (r *queryResolver) GetFollowRequests(ctx context.Context, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetFollowRequests(ctx, input)
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

type connectionResolver struct{ *Resolver }
