package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/util"
)

type CreditTransactionTemplates struct {
	Queries *db.Queries
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByID(ctx context.Context, id uuid.UUID) (*db.CreditTransactionTemplate, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionTemplateByID(ctx, id))
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByName(ctx context.Context, name db.CreditTransactionTemplateName) (*db.CreditTransactionTemplate, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionTemplateByName(ctx, name))
}
