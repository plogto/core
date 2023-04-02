package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/middleware"
)

type DescriptionVariable struct {
	Content string
	Url     *string
	Image   *string
}

func (s *Service) GetCreditTransactionInfoByID(ctx context.Context, id uuid.UUID) (*db.CreditTransactionInfo, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactionInfos.GetCreditTransactionInfoByID(ctx, id)
}
