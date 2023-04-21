package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type TicketMessageAttachments struct {
	Queries *db.Queries
}

func (t *TicketMessageAttachments) CreateTicketMessageAttachment(ctx context.Context, ticketMessageID, fileID pgtype.UUID) (*db.TicketMessageAttachment, error) {
	ticketMessageAttachment, _ := t.Queries.CreateTicketMessageAttachment(ctx, db.CreateTicketMessageAttachmentParams{
		TicketMessageID: ticketMessageID,
		FileID:          fileID,
	})

	return ticketMessageAttachment, nil
}
