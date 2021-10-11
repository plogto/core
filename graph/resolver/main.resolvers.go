package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/plog-core/graph/generated"
	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
)

func (r *mutationResolver) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	fmt.Println(util.RandomString(20))
	return &model.Test{
		Content: input.Content,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
