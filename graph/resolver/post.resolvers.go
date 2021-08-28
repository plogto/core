package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/poster-core/graph/generated"
	"github.com/favecode/poster-core/graph/model"
)

func (r *mutationResolver) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	return r.Service.AddPost(ctx, input)
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

func (r *queryResolver) GetUserPostsByUsername(ctx context.Context, username string, input *model.GetUserPostsByUsernameInput) (*model.Posts, error) {
	return r.Service.GetUserPostsByUsername(ctx, username, input)
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
