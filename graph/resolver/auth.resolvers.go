package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/note-core/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	return r.Service.Register(ctx, input)
}

func (r *queryResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	return r.Service.Login(ctx, input)
}
