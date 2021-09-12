package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/poster-core/graph/generated"
	"github.com/favecode/poster-core/graph/model"
)

func (r *queryResolver) GetTrands(ctx context.Context, input *model.PaginationInput) (*model.Tags, error) {
	return r.Service.GetTrends(ctx, input)
}

func (r *tagResolver) Count(ctx context.Context, obj *model.Tag) (*int, error) {
	return r.Service.CountTagByTagId(ctx, obj.ID)
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
