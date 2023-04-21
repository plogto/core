package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type CreditTransactionTemplates struct {
	Queries *db.Queries
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByID(ctx context.Context, id pgtype.UUID) (*db.CreditTransactionTemplate, error) {
	return c.Queries.GetCreditTransactionTemplateByID(ctx, id)
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByName(ctx context.Context, name db.CreditTransactionTemplateName) (*db.CreditTransactionTemplate, error) {
	return c.Queries.GetCreditTransactionTemplateByName(ctx, name)
}
