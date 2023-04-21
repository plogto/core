package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type CreditTransactionInfos struct {
	Queries *db.Queries
}

func (c *CreditTransactionInfos) CreateCreditTransactionInfo(ctx context.Context, arg db.CreateCreditTransactionInfoParams) (*db.CreditTransactionInfo, error) {
	return c.Queries.CreateCreditTransactionInfo(ctx, arg)
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByID(ctx context.Context, id pgtype.UUID) (*db.CreditTransactionInfo, error) {
	return c.Queries.GetCreditTransactionInfoByID(ctx, id)
}

func (c *CreditTransactionInfos) UpdateCreditTransactionInfoStatus(ctx context.Context, arg db.UpdateCreditTransactionInfoStatusParams) (*db.CreditTransactionInfo, error) {
	return c.Queries.UpdateCreditTransactionInfoStatus(ctx, arg)
}
