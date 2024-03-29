// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: ticket_message_attachments.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTicketMessageAttachment = `-- name: CreateTicketMessageAttachment :one
INSERT INTO
	ticket_message_attachments (ticket_message_id, file_id)
VALUES
	($1, $2) RETURNING id, ticket_message_id, file_id, created_at, deleted_at
`

type CreateTicketMessageAttachmentParams struct {
	TicketMessageID pgtype.UUID
	FileID          pgtype.UUID
}

func (q *Queries) CreateTicketMessageAttachment(ctx context.Context, arg CreateTicketMessageAttachmentParams) (*TicketMessageAttachment, error) {
	row := q.db.QueryRow(ctx, createTicketMessageAttachment, arg.TicketMessageID, arg.FileID)
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
