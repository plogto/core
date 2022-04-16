package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUserInput) (*model.User, error) {
	return r.Service.EditUser(ctx, input)
}

func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.AuthResponse, error) {
	return r.Service.ChangePassword(ctx, input)
}

func (r *queryResolver) GetUserInfo(ctx context.Context) (*model.User, error) {
	return r.Service.GetUserInfo(ctx)
}

func (r *queryResolver) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return r.Service.GetUserByUsername(ctx, username)
}

func (r *queryResolver) CheckUsername(ctx context.Context, username string) (*model.User, error) {
	return r.Service.CheckUsername(ctx, username)
}

func (r *queryResolver) CheckEmail(ctx context.Context, email string) (*model.User, error) {
	return r.Service.CheckEmail(ctx, email)
}

func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*model.File, error) {
	fmt.Println("avatar", obj.Avatar)
	if obj.Avatar != nil {
		return r.Service.GetFileByFileId(ctx, *obj.Avatar)
	} else {
		return nil, nil
	}
}

func (r *userResolver) Background(ctx context.Context, obj *model.User) (*model.File, error) {
	if obj.Background != nil {
		return r.Service.GetFileByFileId(ctx, *obj.Background)
	} else {
		return nil, nil
	}
}

func (r *userResolver) ConnectionStatus(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionStatus(ctx, obj.ID)
}

func (r *userResolver) FollowingCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "following")
}

func (r *userResolver) FollowersCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "followers")
}

func (r *userResolver) FollowRequestsCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionCount(ctx, obj.ID, "requests")
}

func (r *userResolver) PostsCount(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetPostsCount(ctx, obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
