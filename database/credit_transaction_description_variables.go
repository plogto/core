package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type CreditTransactionDescriptionVariables struct {
	Queries *db.Queries
}

func (c *CreditTransactionDescriptionVariables) CreateCreditTransactionDescriptionVariable(ctx context.Context, arg db.CreateCreditTransactionDescriptionVariableParams) (*db.CreditTransactionDescriptionVariable, error) {
	return util.HandleDBResponse(c.Queries.CreateCreditTransactionDescriptionVariable(ctx, arg))

}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByContentID(ctx context.Context, contentID uuid.UUID) (*db.CreditTransactionDescriptionVariable, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionDescriptionVariableByContentID(ctx, contentID))

}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID uuid.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariables, _ := c.Queries.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)

	return creditTransactionDescriptionVariables, nil
}
