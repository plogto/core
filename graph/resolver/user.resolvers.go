package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/poster-core/graph/generated"
	"github.com/favecode/poster-core/graph/model"
)

func (r *queryResolver) GetUserInfo(ctx context.Context) (*model.User, error) {
	return r.Service.GetUserInfo(ctx)
}

func (r *queryResolver) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return r.Service.GetUserByUsername(ctx, username)
}

func (r *userResolver) ConnectionStatus(ctx context.Context, obj *model.User) (*int, error) {
	return r.Service.GetConnectionStatus(ctx, obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
