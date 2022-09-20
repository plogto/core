package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// CreateTicket is the resolver for the createTicket field.
func (r *mutationResolver) CreateTicket(ctx context.Context, input model.CreateTicketInput) (*model.Ticket, error) {
	return r.Service.CreateTicket(ctx, input)
}

// AddTicketMessage is the resolver for the addTicketMessage field.
func (r *mutationResolver) AddTicketMessage(ctx context.Context, ticketID string, input model.AddTicketMessageInput) (*model.TicketMessage, error) {
	return r.Service.AddTicketMessage(ctx, ticketID, input)
}

// ReadTicketMessages is the resolver for the readTicketMessages field.
func (r *mutationResolver) ReadTicketMessages(ctx context.Context, ticketID string) (*model.TicketMessages, error) {
	panic(fmt.Errorf("not implemented: ReadTicketMessages - readTicketMessages"))
}

// GetTickets is the resolver for the getTickets field.
func (r *queryResolver) GetTickets(ctx context.Context, pageInfo *model.PageInfoInput) (*model.Tickets, error) {
	return r.Service.GetTickets(ctx, pageInfo)
}

// GetTicketMessagesByTicketURL is the resolver for the getTicketMessagesByTicketUrl field.
func (r *queryResolver) GetTicketMessagesByTicketURL(ctx context.Context, ticketURL string, pageInfo *model.PageInfoInput) (*model.TicketMessages, error) {
	return r.Service.GetTicketMessagesByTicketURL(ctx, ticketURL, pageInfo)
}

// User is the resolver for the user field.
func (r *ticketResolver) User(ctx context.Context, obj *model.Ticket) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// LastMessage is the resolver for the lastMessage field.
func (r *ticketResolver) LastMessage(ctx context.Context, obj *model.Ticket) (*model.TicketMessage, error) {
	return r.Service.GetLastTicketMessageByTicketID(ctx, obj.ID)
}

// Sender is the resolver for the sender field.
func (r *ticketMessageResolver) Sender(ctx context.Context, obj *model.TicketMessage) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.SenderID)
}

// Ticket is the resolver for the ticket field.
func (r *ticketMessageResolver) Ticket(ctx context.Context, obj *model.TicketMessage) (*model.Ticket, error) {
	return r.Service.GetTicketByID(ctx, obj.TicketID)
}

// Attachment is the resolver for the attachment field.
func (r *ticketMessageResolver) Attachment(ctx context.Context, obj *model.TicketMessage) ([]*model.File, error) {
	return r.Service.GetTicketMessageAttachmentsByTicketMessageID(ctx, obj.ID)
}

// Ticket is the resolver for the ticket field.
func (r *ticketMessagesResolver) Ticket(ctx context.Context, obj *model.TicketMessages) (*model.Ticket, error) {
	return r.Service.GetTicketByID(ctx, obj.Edges[0].Node.TicketID)
}

// Cursor is the resolver for the cursor field.
func (r *ticketMessagesEdgeResolver) Cursor(ctx context.Context, obj *model.TicketMessagesEdge) (string, error) {
	return util.ConvertCreateAtToCursor(*obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *ticketMessagesEdgeResolver) Node(ctx context.Context, obj *model.TicketMessagesEdge) (*model.TicketMessage, error) {
	return r.Service.GetTicketMessageByID(ctx, obj.Node.ID)
}

// Cursor is the resolver for the cursor field.
func (r *ticketsEdgeResolver) Cursor(ctx context.Context, obj *model.TicketsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(*obj.Node.UpdatedAt), nil
}

// Node is the resolver for the node field.
func (r *ticketsEdgeResolver) Node(ctx context.Context, obj *model.TicketsEdge) (*model.Ticket, error) {
	return r.Service.GetTicketByID(ctx, obj.Node.ID)
}

// Ticket returns generated.TicketResolver implementation.
func (r *Resolver) Ticket() generated.TicketResolver { return &ticketResolver{r} }

// TicketMessage returns generated.TicketMessageResolver implementation.
func (r *Resolver) TicketMessage() generated.TicketMessageResolver { return &ticketMessageResolver{r} }

// TicketMessages returns generated.TicketMessagesResolver implementation.
func (r *Resolver) TicketMessages() generated.TicketMessagesResolver {
	return &ticketMessagesResolver{r}
}

// TicketMessagesEdge returns generated.TicketMessagesEdgeResolver implementation.
func (r *Resolver) TicketMessagesEdge() generated.TicketMessagesEdgeResolver {
	return &ticketMessagesEdgeResolver{r}
}

// TicketsEdge returns generated.TicketsEdgeResolver implementation.
func (r *Resolver) TicketsEdge() generated.TicketsEdgeResolver { return &ticketsEdgeResolver{r} }

type ticketResolver struct{ *Resolver }
type ticketMessageResolver struct{ *Resolver }
type ticketMessagesResolver struct{ *Resolver }
type ticketMessagesEdgeResolver struct{ *Resolver }
type ticketsEdgeResolver struct{ *Resolver }