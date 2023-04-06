package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

func (s *Service) GetTicketMessageAttachmentsByTicketMessageID(ctx context.Context, ticketMessageID uuid.UUID) ([]*db.File, error) {
	ticketMessageAttachments, _ := s.Files.GetFilesByTicketMessageID(ctx, ticketMessageID)

	return ticketMessageAttachments, nil
}
