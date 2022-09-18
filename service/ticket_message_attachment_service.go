package service

import (
	"context"

	"github.com/plogto/core/graph/model"
)

func (s *Service) GetTicketMessageAttachmentsByTicketMessageID(ctx context.Context, ticketMessageID string) ([]*model.File, error) {
	ticketMessageAttachments, _ := s.Files.GetFilesByTicketMessageID(ticketMessageID)

	return ticketMessageAttachments, nil
}
