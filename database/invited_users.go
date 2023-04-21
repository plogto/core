package database

import (
	"context"

	"github.com/plogto/core/db"
)

type InvitedUsers struct {
	Queries *db.Queries
}

func (i *InvitedUsers) CreateInvitedUser(ctx context.Context, arg db.CreateInvitedUserParams) (*db.InvitedUser, error) {
	return i.Queries.CreateInvitedUser(ctx, arg)
}
