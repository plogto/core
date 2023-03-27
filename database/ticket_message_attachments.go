package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type TicketMessageAttachments struct {
	Queries *db.Queries
}

func (t *TicketMessageAttachments) CreateTicketMessageAttachment(ctx context.Context, ticketMessageID, fileID uuid.UUID) (*db.TicketMessageAttachment, error) {
	ticketMessageAttachment, err := t.Queries.CreateTicketMessageAttachment(ctx, db.CreateTicketMessageAttachmentParams{
		TicketMessageID: ticketMessageID,
		FileID:          fileID,
	})

	if err != nil {
		return nil, err
	}

	return ticketMessageAttachment, nil
}
