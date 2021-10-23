package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
)

func (r *commentResolver) Parent(ctx context.Context, obj *model.Comment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Children(ctx context.Context, obj *model.Comment) (*model.Comments, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Post(ctx context.Context, obj *model.Comment) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddComment(ctx context.Context, input model.CommentPostInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetComments(ctx context.Context, postID string, input *model.PaginationInput) (*model.Comments, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
