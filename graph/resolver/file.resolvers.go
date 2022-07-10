package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/plogto/core/graph/model"
)

// SingleUploadFile is the resolver for the singleUploadFile field.
func (r *mutationResolver) SingleUploadFile(ctx context.Context, file graphql.Upload) (*model.File, error) {
	return r.Service.SingleUploadFile(ctx, file)
}
