package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

func (s *Service) GetTicketMessageAttachmentsByTicketMessageID(ctx context.Context, ticketMessageID pgtype.UUID) ([]*db.File, error) {
	ticketMessageAttachments, _ := s.Files.GetFilesByTicketMessageID(ctx, ticketMessageID)

	return ticketMessageAttachments, nil
}
