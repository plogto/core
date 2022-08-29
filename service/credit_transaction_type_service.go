package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionTypeByID(ctx context.Context, id *string) (*model.CreditTransactionType, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactionTypes.GetCreditTransactionTypeByID(*id)
}
