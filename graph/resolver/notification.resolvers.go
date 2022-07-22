package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
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

// Cursor is the resolver for the cursor field.
func (r *notificationsEdgeResolver) Cursor(ctx context.Context, obj *model.NotificationsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(*obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *notificationsEdgeResolver) Node(ctx context.Context, obj *model.NotificationsEdge) (*model.Notification, error) {
	return r.Service.GetNotificationByID(ctx, obj.Node.ID)
}

// GetNotifications is the resolver for the getNotifications field.
func (r *queryResolver) GetNotifications(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.Notifications, error) {
	return r.Service.GetNotifications(ctx, pageInfoInput)
}

// GetNotification is the resolver for the getNotification field.
func (r *subscriptionResolver) GetNotification(ctx context.Context) (<-chan *model.Notification, error) {
	return r.Service.GetNotification(ctx)
}

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

// NotificationsEdge returns generated.NotificationsEdgeResolver implementation.
func (r *Resolver) NotificationsEdge() generated.NotificationsEdgeResolver {
	return &notificationsEdgeResolver{r}
}

type notificationResolver struct{ *Resolver }
type notificationsEdgeResolver struct{ *Resolver }
