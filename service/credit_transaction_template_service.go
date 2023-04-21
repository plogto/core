package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionTemplateByID(ctx context.Context, id pgtype.UUID) (*db.CreditTransactionTemplate, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactionTemplates.GetCreditTransactionTemplateByID(ctx, id)
}
