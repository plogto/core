package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// EditUser is the resolver for the editUser field.
func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUserInput) (*model.User, error) {
	return r.Service.EditUser(ctx, input)
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.AuthResponse, error) {
	return r.Service.ChangePassword(ctx, input)
}

// GetUserInfo is the resolver for the getUserInfo field.
func (r *queryResolver) GetUserInfo(ctx context.Context) (*model.User, error) {
	return r.Service.GetUserInfo(ctx)
}

// GetUserByUsername is the resolver for the getUserByUsername field.
func (r *queryResolver) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return r.Service.GetUserByUsername(ctx, username)
}

// CheckUsername is the resolver for the checkUsername field.
func (r *queryResolver) CheckUsername(ctx context.Context, username string) (*model.User, error) {
	return r.Service.CheckUsername(ctx, username)
}

// CheckEmail is the resolver for the checkEmail field.
func (r *queryResolver) CheckEmail(ctx context.Context, email string) (*model.User, error) {
	return r.Service.CheckEmail(ctx, email)
}

// Avatar is the resolver for the avatar field.
func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*model.File, error) {
	if obj.Avatar != nil {
		return r.Service.GetFileByFileId(ctx, *obj.Avatar)
	} else {
		return nil, nil
	}
}

// Background is the resolver for the background field.
func (r *userResolver) Background(ctx context.Context, obj *model.User) (*model.File, error) {
	if obj.Background != nil {
		return r.Service.GetFileByFileId(ctx, *obj.Background)
	} else {
		return nil, nil
	}
}

// ConnectionStatus is the resolver for the connectionStatus field.
func (r *userResolver) ConnectionStatus(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionStatus(ctx, obj.ID)
}

// FollowingCount is the resolver for the followingCount field.
func (r *userResolver) FollowingCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "following")
}

// FollowersCount is the resolver for the followersCount field.
func (r *userResolver) FollowersCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "followers")
}

// FollowRequestsCount is the resolver for the followRequestsCount field.
func (r *userResolver) FollowRequestsCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "requests")
}

// PostsCount is the resolver for the postsCount field.
func (r *userResolver) PostsCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetPostsCount(ctx, obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) IsVerified(ctx context.Context, obj *model.User) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
