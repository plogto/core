// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: ticket_message_attachments.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTicketMessageAttachment = `-- name: CreateTicketMessageAttachment :one
INSERT INTO
	ticket_message_attachments (ticket_message_id, file_id)
VALUES
	($1, $2) RETURNING id, ticket_message_id, file_id, created_at, deleted_at
`

type CreateTicketMessageAttachmentParams struct {
	TicketMessageID uuid.UUID
	FileID          uuid.UUID
}

func (q *Queries) CreateTicketMessageAttachment(ctx context.Context, arg CreateTicketMessageAttachmentParams) (*TicketMessageAttachment, error) {
	row := q.db.QueryRowContext(ctx, createTicketMessageAttachment, arg.TicketMessageID, arg.FileID)
	var i TicketMessageAttachment
	err := row.Scan(
		&i.ID,
		&i.TicketMessageID,
		&i.FileID,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
