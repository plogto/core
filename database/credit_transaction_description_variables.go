package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type CreditTransactionDescriptionVariables struct {
	Queries *db.Queries
}

func (c *CreditTransactionDescriptionVariables) CreateCreditTransactionDescriptionVariable(ctx context.Context, arg db.CreateCreditTransactionDescriptionVariableParams) (*db.CreditTransactionDescriptionVariable, error) {
	return c.Queries.CreateCreditTransactionDescriptionVariable(ctx, arg)

}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByContentID(ctx context.Context, contentID pgtype.UUID) (*db.CreditTransactionDescriptionVariable, error) {
	return c.Queries.GetCreditTransactionDescriptionVariableByContentID(ctx, contentID)

}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID pgtype.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	creditTransactionDescriptionVariables, _ := c.Queries.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)

	return creditTransactionDescriptionVariables, nil
}
