package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, expression string) (*model.Search, error) {
	return r.Service.Search(ctx, expression)
}
