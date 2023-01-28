package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type CreditTransactionTemplates struct {
	Queries *db.Queries
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByID(ctx context.Context, id uuid.UUID) (*db.CreditTransactionTemplate, error) {
	creditTransactionTemplate, err := c.Queries.GetCreditTransactionTemplateByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return creditTransactionTemplate, err
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByName(ctx context.Context, name db.CreditTransactionTemplateName) (*db.CreditTransactionTemplate, error) {
	creditTransactionTemplate, err := c.Queries.GetCreditTransactionTemplateByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return creditTransactionTemplate, err

}
