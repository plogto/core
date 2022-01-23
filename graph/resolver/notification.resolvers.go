package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

func (r *notificationResolver) NotificationType(ctx context.Context, obj *model.Notification) (*model.NotificationType, error) {
	return r.Service.GetNotificationType(ctx, obj.NotificationTypeID)
}

func (r *notificationResolver) Sender(ctx context.Context, obj *model.Notification) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.SenderID)
}

func (r *notificationResolver) Receiver(ctx context.Context, obj *model.Notification) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.ReceiverID)
}

func (r *notificationResolver) Post(ctx context.Context, obj *model.Notification) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.PostID)
}

func (r *notificationResolver) Comment(ctx context.Context, obj *model.Notification) (*model.Comment, error) {
	return r.Service.GetCommentByID(ctx, obj.CommentID)
}

func (r *queryResolver) GetNotifications(ctx context.Context, input *model.PaginationInput) (*model.Notifications, error) {
	return r.Service.GetNotifications(ctx, input)
}

func (r *subscriptionResolver) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Service.GetNotification(ctx)
}

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

type notificationResolver struct{ *Resolver }
