package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// GetTagByTagName is the resolver for the getTagByTagName field.
func (r *queryResolver) GetTagByTagName(ctx context.Context, tagName string) (*model.Tag, error) {
	return r.Service.GetTagByName(ctx, tagName)
}

// GetTrends is the resolver for the getTrends field.
func (r *queryResolver) GetTrends(ctx context.Context, first *int) (*model.Tags, error) {
	return r.Service.GetTrends(ctx, first)
}

// Node is the resolver for the node field.
func (r *tagsEdgeResolver) Node(ctx context.Context, obj *model.TagsEdge) (*model.Tag, error) {
	return r.Service.GetTagByID(ctx, obj.Node.ID)
}

// TagsEdge returns generated.TagsEdgeResolver implementation.
func (r *Resolver) TagsEdge() generated.TagsEdgeResolver { return &tagsEdgeResolver{r} }

type tagsEdgeResolver struct{ *Resolver }
