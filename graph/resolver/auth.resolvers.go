package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/model"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	return r.Service.Register(ctx, input, false)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	return r.Service.Login(ctx, input)
}

// OAuthGoogle is the resolver for the oAuthGoogle field.
func (r *queryResolver) OAuthGoogle(ctx context.Context, input model.OAuthGoogleInput) (*model.AuthResponse, error) {
	return r.Service.OAuthGoogle(ctx, input)
}
