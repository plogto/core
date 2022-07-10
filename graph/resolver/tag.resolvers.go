package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/model"
)

// GetTagByTagName is the resolver for the getTagByTagName field.
func (r *queryResolver) GetTagByTagName(ctx context.Context, tagName string) (*model.Tag, error) {
	return r.Service.GetTagByName(ctx, tagName)
}

// GetTrends is the resolver for the getTrends field.
func (r *queryResolver) GetTrends(ctx context.Context, input *model.PaginationInput) (*model.Tags, error) {
	return r.Service.GetTrends(ctx, input)
}
