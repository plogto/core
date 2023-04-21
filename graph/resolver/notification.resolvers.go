package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// ReadNotifications is the resolver for the readNotifications field.
func (r *mutationResolver) ReadNotifications(ctx context.Context) (*bool, error) {
	return r.Service.ReadNotifications(ctx)
}

// NotificationType is the resolver for the notificationType field.
func (r *notificationResolver) NotificationType(ctx context.Context, obj *db.Notification) (*db.NotificationType, error) {
	return r.Service.GetNotificationType(ctx, obj.NotificationTypeID)
}

// Sender is the resolver for the sender field.
func (r *notificationResolver) Sender(ctx context.Context, obj *db.Notification) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.SenderID)
}

// Receiver is the resolver for the receiver field.
func (r *notificationResolver) Receiver(ctx context.Context, obj *db.Notification) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.ReceiverID)
}

// Post is the resolver for the post field.
func (r *notificationResolver) Post(ctx context.Context, obj *db.Notification) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.PostID)
}

// Reply is the resolver for the reply field.
func (r *notificationResolver) Reply(ctx context.Context, obj *db.Notification) (*model.Post, error) {
	return r.Service.GetPostByID(ctx, obj.ReplyID)
}

// Name is the resolver for the name field.
func (r *notificationTypeResolver) Name(ctx context.Context, obj *db.NotificationType) (model.NotificationTypeName, error) {
	return model.NotificationTypeName(obj.Name), nil
}

// Cursor is the resolver for the cursor field.
func (r *notificationsEdgeResolver) Cursor(ctx context.Context, obj *model.NotificationsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *notificationsEdgeResolver) Node(ctx context.Context, obj *model.NotificationsEdge) (*db.Notification, error) {
	return r.Service.GetNotificationByID(ctx, obj.Node.ID)
}

// GetNotifications is the resolver for the getNotifications field.
func (r *queryResolver) GetNotifications(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Notifications, error) {
	return r.Service.GetNotifications(ctx, pageInfo)
}

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

// NotificationType returns generated.NotificationTypeResolver implementation.
func (r *Resolver) NotificationType() generated.NotificationTypeResolver {
	return &notificationTypeResolver{r}
}

// NotificationsEdge returns generated.NotificationsEdgeResolver implementation.
func (r *Resolver) NotificationsEdge() generated.NotificationsEdgeResolver {
	return &notificationsEdgeResolver{r}
}

type notificationResolver struct{ *Resolver }
type notificationTypeResolver struct{ *Resolver }
type notificationsEdgeResolver struct{ *Resolver }
