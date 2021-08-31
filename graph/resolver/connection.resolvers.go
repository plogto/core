package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/poster-core/graph/generated"
	"github.com/favecode/poster-core/graph/model"
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

func (r *queryResolver) GetUserFollowersByUsername(ctx context.Context, username string, input *model.GetUserConnectionsByUserIDInput) (*model.Connections, error) {
	return r.Service.GetUserConnectionsByUsername(ctx, username, input, "follower")
}

func (r *queryResolver) GetUserFollowingByUsername(ctx context.Context, username string, input *model.GetUserConnectionsByUserIDInput) (*model.Connections, error) {
	return r.Service.GetUserConnectionsByUsername(ctx, username, input, "following")
}

func (r *queryResolver) GetUserFollowRequests(ctx context.Context, input *model.GetUserConnectionsByUserIDInput) (*model.Connections, error) {
	return r.Service.GetUserFollowRequests(ctx, input)
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

type connectionResolver struct{ *Resolver }
