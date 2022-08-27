package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateCreditTransactionDescriptionVariableInput struct {
	SenderID    string
	ReceiverID  string
	Amount      float64
	Description *string
	Status      model.CreditTransactionStatus
}

func (s *Service) GetCreditTransactionDescriptionVariableByID(ctx context.Context, id *string) (*model.CreditTransactionDescriptionVariable, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariableByID(*id)
}

func (s *Service) GetCreditTransactionDescriptionVariables(ctx context.Context, input *model.PageInfoInput) (*model.CreditTransactions, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)

	return s.CreditTransactions.GetCreditTransactionsByUserIDAndPageInfo(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}
