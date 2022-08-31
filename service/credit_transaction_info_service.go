package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
)

type DescriptionVariable struct {
	Content string
	Url     *string
	Image   *string
}

func (s *Service) CreateCreditTransactionInfo(creditTransactionInfo model.CreditTransactionInfo) (*model.CreditTransactionInfo, error) {
	return s.CreditTransactionInfos.CreateCreditTransactionInfo(&creditTransactionInfo)
}

func (s *Service) GetCreditTransactionInfoByID(ctx context.Context, id *string) (*model.CreditTransactionInfo, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactionInfos.GetCreditTransactionInfoByID(*id)
}
