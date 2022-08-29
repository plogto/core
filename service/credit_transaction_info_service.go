package service

import (
	"github.com/plogto/core/graph/model"
)

type DescriptionVariable struct {
	Content string
	Url     *string
	Image   *string
}

func (s *Service) CreateCreditTransactionInfo(creditTransactionInfo model.CreditTransactionInfo) (*model.CreditTransactionInfo, error) {
	return s.CreditTransactionInfos.CreateCreditTransactionInfo(&creditTransactionInfo)
}
