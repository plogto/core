package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
)

// Inviter is the resolver for the inviter field.
func (r *invitedUserResolver) Inviter(ctx context.Context, obj *model.InvitedUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Inviter - inviter"))
}

// Invitee is the resolver for the invitee field.
func (r *invitedUserResolver) Invitee(ctx context.Context, obj *model.InvitedUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Invitee - invitee"))
}

// Cursor is the resolver for the cursor field.
func (r *invitedUsersEdgeResolver) Cursor(ctx context.Context, obj *model.InvitedUsersEdge) (string, error) {
	panic(fmt.Errorf("not implemented: Cursor - cursor"))
}

// Node is the resolver for the node field.
func (r *invitedUsersEdgeResolver) Node(ctx context.Context, obj *model.InvitedUsersEdge) (*model.InvitedUser, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// GetInvitedUsers is the resolver for the getInvitedUsers field.
func (r *queryResolver) GetInvitedUsers(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.InvitedUsers, error) {
	panic(fmt.Errorf("not implemented: GetInvitedUsers - getInvitedUsers"))
}

// InvitedUser returns generated.InvitedUserResolver implementation.
func (r *Resolver) InvitedUser() generated.InvitedUserResolver { return &invitedUserResolver{r} }

// InvitedUsersEdge returns generated.InvitedUsersEdgeResolver implementation.
func (r *Resolver) InvitedUsersEdge() generated.InvitedUsersEdgeResolver {
	return &invitedUsersEdgeResolver{r}
}

type invitedUserResolver struct{ *Resolver }
type invitedUsersEdgeResolver struct{ *Resolver }
