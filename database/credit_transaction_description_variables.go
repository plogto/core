package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type CreditTransactionDescriptionVariables struct {
	Queries *db.Queries
}

func (c *CreditTransactionDescriptionVariables) CreateCreditTransactionDescriptionVariable(ctx context.Context, arg db.CreateCreditTransactionDescriptionVariableParams) (*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, _ := c.Queries.CreateCreditTransactionDescriptionVariable(ctx, arg)

	return creditTransactionDescriptionVariable, nil
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByContentID(ctx context.Context, contentID uuid.UUID) (*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, _ := c.Queries.GetCreditTransactionDescriptionVariableByContentID(ctx, contentID)

	return creditTransactionDescriptionVariable, nil
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID uuid.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariable, _ := c.Queries.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)

	return creditTransactionDescriptionVariable, nil
}
