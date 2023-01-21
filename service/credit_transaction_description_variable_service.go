package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID uuid.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	// FIXME
	if err != nil {
		return nil, nil
	}

	return s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)
}
