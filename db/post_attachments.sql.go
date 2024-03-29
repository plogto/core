// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: post_attachments.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPostAttachment = `-- name: CreatePostAttachment :one
INSERT INTO
	post_attachments (post_id, file_id)
VALUES
	($1, $2) RETURNING id, post_id, file_id, created_at, deleted_at
`

type CreatePostAttachmentParams struct {
	PostID pgtype.UUID
	FileID pgtype.UUID
}

func (q *Queries) CreatePostAttachment(ctx context.Context, arg CreatePostAttachmentParams) (*PostAttachment, error) {
	row := q.db.QueryRow(ctx, createPostAttachment, arg.PostID, arg.FileID)
	var i PostAttachment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.FileID,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
