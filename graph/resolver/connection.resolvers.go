package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *connectionResolver) Following(ctx context.Context, obj *model.Connection) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowingID)
}

func (r *connectionResolver) Follower(ctx context.Context, obj *model.Connection) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowerID)
}

func (r *mutationResolver) FollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.FollowUser(ctx, userID)
}

func (r *mutationResolver) UnfollowUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.UnfollowUser(ctx, userID)
}

func (r *mutationResolver) AcceptUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.AcceptUser(ctx, userID)
}

func (r *mutationResolver) RejectUser(ctx context.Context, userID string) (*model.Connection, error) {
	return r.Service.RejectUser(ctx, userID)
}

func (r *queryResolver) GetFollowersByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, input, "followers")
}

func (r *queryResolver) GetFollowingByUsername(ctx context.Context, username string, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, input, "following")
}

func (r *queryResolver) GetFollowRequests(ctx context.Context, input *model.PaginationInput) (*model.Connections, error) {
	return r.Service.GetFollowRequests(ctx, input)
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

type connectionResolver struct{ *Resolver }
