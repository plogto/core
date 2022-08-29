package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionTemplateByID(ctx context.Context, id *string) (*model.CreditTransactionTemplate, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactionTemplates.GetCreditTransactionTemplateByID(*id)
}
