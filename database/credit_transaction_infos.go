package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type CreditTransactionInfos struct {
	Queries *db.Queries
}

func (c *CreditTransactionInfos) CreateCreditTransactionInfo(ctx context.Context, arg db.CreateCreditTransactionInfoParams) (*db.CreditTransactionInfo, error) {
	creditTransactionInfo, _ := c.Queries.CreateCreditTransactionInfo(ctx, arg)

	return creditTransactionInfo, nil
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByID(ctx context.Context, id uuid.UUID) (*db.CreditTransactionInfo, error) {
	creditTransactionInfo, _ := c.Queries.GetCreditTransactionInfoByID(ctx, id)

	return creditTransactionInfo, nil
}

func (c *CreditTransactionInfos) UpdateCreditTransactionInfoStatus(ctx context.Context, arg db.UpdateCreditTransactionInfoStatusParams) (*db.CreditTransactionInfo, error) {
	creditTransactionInfo, _ := c.Queries.UpdateCreditTransactionInfoStatus(ctx, arg)

	return creditTransactionInfo, nil
}
