package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

type CreateCreditTransactionDescriptionVariableInput struct {
	SenderID    string
	ReceiverID  string
	Amount      float64
	Description *string
	Status      model.CreditTransactionStatus
}

func (s *Service) GetCreditTransactionDescriptionVariablesByCreditTransactionID(ctx context.Context, creditTransactionID *string) ([]*model.CreditTransactionDescriptionVariable, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if creditTransactionID == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariablesByCreditTransactionID(*creditTransactionID)
}
