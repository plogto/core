package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionTemplateByID(ctx context.Context, id uuid.NullUUID) (*db.CreditTransactionTemplate, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactionTemplates.GetCreditTransactionTemplateByID(ctx, id.UUID)
}
