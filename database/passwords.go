package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type Passwords struct {
	Queries *db.Queries
}

func (p *Passwords) GetPasswordByUserID(ctx context.Context, id pgtype.UUID) (*db.Password, error) {
	return p.Queries.GetPasswordByUserID(ctx, id)
}

func (p *Passwords) AddPassword(ctx context.Context, arg db.CreatePasswordParams) (*db.Password, error) {
	return p.Queries.CreatePassword(ctx, arg)
}

func (p *Passwords) UpdatePassword(ctx context.Context, arg db.UpdatePasswordParams) (*db.Password, error) {
	return p.Queries.UpdatePassword(ctx, arg)
}
