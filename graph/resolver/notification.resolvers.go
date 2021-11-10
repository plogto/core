package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/plog-core/graph/model"
)

func (r *subscriptionResolver) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Service.GetNotification(ctx)
}
