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
	creditTransactionInfo, err := c.Queries.CreateCreditTransactionInfo(ctx, arg)

	if err != nil {
		return nil, err
	}

	return creditTransactionInfo, err
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByID(ctx context.Context, id string) (*db.CreditTransactionInfo, error) {
	ID, _ := uuid.Parse(id)
	creditTransactionInfo, err := c.Queries.GetCreditTransactionInfoByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return creditTransactionInfo, err
}

func (c *CreditTransactionInfos) UpdateCreditTransactionInfoStatus(ctx context.Context, arg db.UpdateCreditTransactionInfoStatusParams) (*db.CreditTransactionInfo, error) {
	creditTransactionInfo, err := c.Queries.UpdateCreditTransactionInfoStatus(ctx, arg)

	if err != nil {
		return nil, err
	}

	return creditTransactionInfo, err
}
