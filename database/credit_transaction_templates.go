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
	creditTransactionTemplate, _ := c.Queries.GetCreditTransactionTemplateByID(ctx, id)

	return creditTransactionTemplate, nil
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByName(ctx context.Context, name db.CreditTransactionTemplateName) (*db.CreditTransactionTemplate, error) {
	creditTransactionTemplate, _ := c.Queries.GetCreditTransactionTemplateByName(ctx, name)

	return creditTransactionTemplate, nil

}
