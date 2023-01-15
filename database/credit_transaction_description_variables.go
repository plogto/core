package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type CreditTransactionDescriptionVariables struct {
	Queries *db.Queries
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByContentID(ctx context.Context, contentID uuid.UUID) (*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, err := c.Queries.GetCreditTransactionDescriptionVariableByContentID(ctx, contentID)

	if err != nil {
		return nil, err
	}

	return creditTransactionDescriptionVariable, nil
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID uuid.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, err := c.Queries.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)

	if err != nil {
		return nil, err
	}

	return creditTransactionDescriptionVariable, nil
}

func (c *CreditTransactionDescriptionVariables) CreateCreditTransactionDescriptionVariable(ctx context.Context, arg db.CreateCreditTransactionDescriptionVariableParams) (*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, err := c.Queries.CreateCreditTransactionDescriptionVariable(ctx, arg)

	if err != nil {
		return nil, err
	}

	return creditTransactionDescriptionVariable, nil
}
