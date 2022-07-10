package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// NotificationType is the resolver for the notificationType field.
func (r *notificationResolver) NotificationType(ctx context.Context, obj *model.Notification) (*model.NotificationType, error) {
	return r.Service.GetNotificationType(ctx, obj.NotificationTypeID)
}

// Sender is the resolver for the sender field.
func (r *notificationResolver) Sender(ctx context.Context, obj *model.Notification) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.SenderID)
}

// Receiver is the resolver for the receiver field.
func (r *notificationResolver) Receiver(ctx context.Context, obj *model.Notification) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.ReceiverID)
}

// Post is the resolver for the post field.
func (r *notificationResolver) Post(ctx context.Context, obj *model.Notification) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.PostID)
}

// Reply is the resolver for the reply field.
func (r *notificationResolver) Reply(ctx context.Context, obj *model.Notification) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.ReplyID)
}

// GetNotifications is the resolver for the getNotifications field.
func (r *queryResolver) GetNotifications(ctx context.Context, input *model.PaginationInput) (*model.Notifications, error) {
	return r.Service.GetNotifications(ctx, input)
}

// GetNotification is the resolver for the getNotification field.
func (r *subscriptionResolver) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Service.GetNotification(ctx)
}

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

type notificationResolver struct{ *Resolver }
