package database

import (
	"context"

	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type InvitedUsers struct {
	Queries *db.Queries
}

func (i *InvitedUsers) CreateInvitedUser(ctx context.Context, arg db.CreateInvitedUserParams) (*db.InvitedUser, error) {
	return util.HandleDBResponse(i.Queries.CreateInvitedUser(ctx, arg))
}
