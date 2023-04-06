package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type CreditTransactionInfos struct {
	Queries *db.Queries
}

func (c *CreditTransactionInfos) CreateCreditTransactionInfo(ctx context.Context, arg db.CreateCreditTransactionInfoParams) (*db.CreditTransactionInfo, error) {
	return util.HandleDBResponse(c.Queries.CreateCreditTransactionInfo(ctx, arg))
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByID(ctx context.Context, id uuid.UUID) (*db.CreditTransactionInfo, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionInfoByID(ctx, id))
}

func (c *CreditTransactionInfos) UpdateCreditTransactionInfoStatus(ctx context.Context, arg db.UpdateCreditTransactionInfoStatusParams) (*db.CreditTransactionInfo, error) {
	return util.HandleDBResponse(c.Queries.UpdateCreditTransactionInfoStatus(ctx, arg))
}
