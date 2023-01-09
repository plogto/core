package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

func (s *Service) GetTicketMessageAttachmentsByTicketMessageID(ctx context.Context, ticketMessageID string) ([]*db.File, error) {
	TicketMessageID, _ := uuid.Parse(ticketMessageID)
	ticketMessageAttachments, _ := s.Files.GetFilesByTicketMessageID(ctx, TicketMessageID)

	return ticketMessageAttachments, nil
}
