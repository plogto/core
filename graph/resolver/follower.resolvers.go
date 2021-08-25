package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/note-core/graph/generated"
	"github.com/favecode/note-core/graph/model"
)

func (r *followerResolver) User(ctx context.Context, obj *model.Follower) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *followerResolver) Follower(ctx context.Context, obj *model.Follower) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowerID)
}

func (r *mutationResolver) FollowUser(ctx context.Context, userID string) (*model.Follower, error) {
	return r.Service.FollowUser(ctx, userID)
}

func (r *mutationResolver) UnfollowUser(ctx context.Context, userID string) (*model.Follower, error) {
	return r.Service.UnfollowUser(ctx, userID)
}

func (r *mutationResolver) AcceptUser(ctx context.Context, userID string) (*model.Follower, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RejectUser(ctx context.Context, userID string) (*model.Follower, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserFollowersByUserID(ctx context.Context, userID string, input *model.GetUserFollowersByUserIDInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Follower returns generated.FollowerResolver implementation.
func (r *Resolver) Follower() generated.FollowerResolver { return &followerResolver{r} }

type followerResolver struct{ *Resolver }
