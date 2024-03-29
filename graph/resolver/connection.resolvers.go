package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// Following is the resolver for the following field.
func (r *connectionResolver) Following(ctx context.Context, obj *db.Connection) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowingID)
}

// Follower is the resolver for the follower field.
func (r *connectionResolver) Follower(ctx context.Context, obj *db.Connection) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.FollowerID)
}

// Cursor is the resolver for the cursor field.
func (r *connectionsEdgeResolver) Cursor(ctx context.Context, obj *model.ConnectionsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *connectionsEdgeResolver) Node(ctx context.Context, obj *model.ConnectionsEdge) (*db.Connection, error) {
	return r.Service.Connections.GetConnectionByID(ctx, obj.Node.ID)
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, userID string) (*db.Connection, error) {
	return r.Service.FollowUser(ctx, convertor.StringToUUID(userID))
}

// UnfollowUser is the resolver for the unfollowUser field.
func (r *mutationResolver) UnfollowUser(ctx context.Context, userID string) (*db.Connection, error) {
	return r.Service.UnfollowUser(ctx, convertor.StringToUUID(userID))
}

// AcceptUser is the resolver for the acceptUser field.
func (r *mutationResolver) AcceptUser(ctx context.Context, userID string) (*db.Connection, error) {
	return r.Service.AcceptUser(ctx, convertor.StringToUUID(userID))
}

// RejectUser is the resolver for the rejectUser field.
func (r *mutationResolver) RejectUser(ctx context.Context, userID string) (*db.Connection, error) {
	return r.Service.RejectUser(ctx, convertor.StringToUUID(userID))
}

// GetFollowersByUsername is the resolver for the getFollowersByUsername field.
func (r *queryResolver) GetFollowersByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, pageInfo, constants.Followers)
}

// GetFollowingByUsername is the resolver for the getFollowingByUsername field.
func (r *queryResolver) GetFollowingByUsername(ctx context.Context, username string, pageInfo *model.PageInfoInput) (*model.Connections, error) {
	return r.Service.GetConnectionsByUsername(ctx, username, pageInfo, constants.Following)
}

// GetFollowRequests is the resolver for the getFollowRequests field.
func (r *queryResolver) GetFollowRequests(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Connections, error) {
	return r.Service.GetFollowRequests(ctx, pageInfo)
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

// ConnectionsEdge returns generated.ConnectionsEdgeResolver implementation.
func (r *Resolver) ConnectionsEdge() generated.ConnectionsEdgeResolver {
	return &connectionsEdgeResolver{r}
}

type connectionResolver struct{ *Resolver }
type connectionsEdgeResolver struct{ *Resolver }
