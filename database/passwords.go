package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type Passwords struct {
	Queries *db.Queries
}

func (p *Passwords) GetPasswordByUserID(ctx context.Context, id uuid.UUID) (*db.Password, error) {
	password, err := p.Queries.GetPasswordByUserID(ctx, id)

	if err != nil {
		return nil, err
	}

	return password, nil
}

func (p *Passwords) AddPassword(ctx context.Context, arg db.CreatePasswordParams) (*db.Password, error) {
	password, err := p.Queries.CreatePassword(ctx, arg)

	if err != nil {
		return nil, err
	}

	return password, nil
}

func (p *Passwords) UpdatePassword(ctx context.Context, arg db.UpdatePasswordParams) (*db.Password, error) {
	password, err := p.Queries.UpdatePassword(ctx, arg)

	if err != nil {
		return nil, err
	}

	return password, nil
}
