package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/middleware"
)

func (s *Service) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID pgtype.UUID) ([]*db.CreditTransactionDescriptionVariable, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, creditTransactionInfoID)
}
