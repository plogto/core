package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type Passwords struct {
	Queries *db.Queries
}

func (p *Passwords) GetPasswordByUserID(ctx context.Context, id uuid.UUID) (*db.Password, error) {
	return util.HandleDBResponse(p.Queries.GetPasswordByUserID(ctx, id))
}

func (p *Passwords) AddPassword(ctx context.Context, arg db.CreatePasswordParams) (*db.Password, error) {
	return util.HandleDBResponse(p.Queries.CreatePassword(ctx, arg))
}

func (p *Passwords) UpdatePassword(ctx context.Context, arg db.UpdatePasswordParams) (*db.Password, error) {
	return util.HandleDBResponse(p.Queries.UpdatePassword(ctx, arg))
}
