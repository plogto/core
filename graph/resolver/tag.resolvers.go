package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *queryResolver) GetTagByTagName(ctx context.Context, tagName string) (*model.Tag, error) {
	return r.Service.GetTagByName(ctx, tagName)
}

func (r *queryResolver) GetTrends(ctx context.Context, input *model.PaginationInput) (*model.Tags, error) {
	return r.Service.GetTrends(ctx, input)
}

func (r *tagResolver) Count(ctx context.Context, obj *model.Tag) (*int, error) {
	return r.Service.CountTagByTagID(ctx, obj.ID)
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
