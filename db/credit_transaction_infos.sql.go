// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: credit_transaction_infos.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCreditTransactionInfo = `-- name: CreateCreditTransactionInfo :one
INSERT INTO
	credit_transaction_infos (
		description,
		credit_transaction_template_id,
		status
	)
VALUES
	($1, $2, $3) RETURNING id, credit_transaction_template_id, description, status, created_at, updated_at, deleted_at
`

type CreateCreditTransactionInfoParams struct {
	Description                 pgtype.Text
	CreditTransactionTemplateID pgtype.UUID
	Status                      CreditTransactionStatus
}

func (q *Queries) CreateCreditTransactionInfo(ctx context.Context, arg CreateCreditTransactionInfoParams) (*CreditTransactionInfo, error) {
	row := q.db.QueryRow(ctx, createCreditTransactionInfo, arg.Description, arg.CreditTransactionTemplateID, arg.Status)
	var i CreditTransactionInfo
	err := row.Scan(
		&i.ID,
		&i.CreditTransactionTemplateID,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getCreditTransactionInfoByID = `-- name: GetCreditTransactionInfoByID :one
SELECT
	id, credit_transaction_template_id, description, status, created_at, updated_at, deleted_at
FROM
	credit_transaction_infos
WHERE
	id = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetCreditTransactionInfoByID(ctx context.Context, id pgtype.UUID) (*CreditTransactionInfo, error) {
	row := q.db.QueryRow(ctx, getCreditTransactionInfoByID, id)
	var i CreditTransactionInfo
	err := row.Scan(
		&i.ID,
		&i.CreditTransactionTemplateID,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateCreditTransactionInfoStatus = `-- name: UpdateCreditTransactionInfoStatus :one
UPDATE
	credit_transaction_infos
SET
	status = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING id, credit_transaction_template_id, description, status, created_at, updated_at, deleted_at
`

type UpdateCreditTransactionInfoStatusParams struct {
	Status CreditTransactionStatus
	ID     pgtype.UUID
}

func (q *Queries) UpdateCreditTransactionInfoStatus(ctx context.Context, arg UpdateCreditTransactionInfoStatusParams) (*CreditTransactionInfo, error) {
	row := q.db.QueryRow(ctx, updateCreditTransactionInfoStatus, arg.Status, arg.ID)
	var i CreditTransactionInfo
	err := row.Scan(
		&i.ID,
		&i.CreditTransactionTemplateID,
		&i.Description,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
