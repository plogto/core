// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: credit_transaction_description_variables.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCreditTransactionDescriptionVariable = `-- name: CreateCreditTransactionDescriptionVariable :one
INSERT INTO
	credit_transaction_description_variables (credit_transaction_info_id, content_id, KEY, TYPE)
VALUES
	($1, $2, $3, $4) RETURNING id, credit_transaction_info_id, content_id, key, type, created_at, deleted_at
`

type CreateCreditTransactionDescriptionVariableParams struct {
	CreditTransactionInfoID pgtype.UUID
	ContentID               pgtype.UUID
	Key                     CreditTransactionDescriptionVariableKey
	Type                    CreditTransactionDescriptionVariableType
}

func (q *Queries) CreateCreditTransactionDescriptionVariable(ctx context.Context, arg CreateCreditTransactionDescriptionVariableParams) (*CreditTransactionDescriptionVariable, error) {
	row := q.db.QueryRow(ctx, createCreditTransactionDescriptionVariable,
		arg.CreditTransactionInfoID,
		arg.ContentID,
		arg.Key,
		arg.Type,
	)
	var i CreditTransactionDescriptionVariable
	err := row.Scan(
		&i.ID,
		&i.CreditTransactionInfoID,
		&i.ContentID,
		&i.Key,
		&i.Type,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getCreditTransactionDescriptionVariableByContentID = `-- name: GetCreditTransactionDescriptionVariableByContentID :one
SELECT
	id, credit_transaction_info_id, content_id, key, type, created_at, deleted_at
FROM
	credit_transaction_description_variables
WHERE
	content_id = $1
	AND deleted_at IS NULL
ORDER BY
	created_at DESC
LIMIT
	1
`

func (q *Queries) GetCreditTransactionDescriptionVariableByContentID(ctx context.Context, contentID pgtype.UUID) (*CreditTransactionDescriptionVariable, error) {
	row := q.db.QueryRow(ctx, getCreditTransactionDescriptionVariableByContentID, contentID)
	var i CreditTransactionDescriptionVariable
	err := row.Scan(
		&i.ID,
		&i.CreditTransactionInfoID,
		&i.ContentID,
		&i.Key,
		&i.Type,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getCreditTransactionDescriptionVariablesByCreditTransactionInfoID = `-- name: GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID :many
SELECT
	id, credit_transaction_info_id, content_id, key, type, created_at, deleted_at
FROM
	credit_transaction_description_variables
WHERE
	credit_transaction_info_id = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx context.Context, creditTransactionInfoID pgtype.UUID) ([]*CreditTransactionDescriptionVariable, error) {
	rows, err := q.db.Query(ctx, getCreditTransactionDescriptionVariablesByCreditTransactionInfoID, creditTransactionInfoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*CreditTransactionDescriptionVariable{}
	for rows.Next() {
		var i CreditTransactionDescriptionVariable
		if err := rows.Scan(
			&i.ID,
			&i.CreditTransactionInfoID,
			&i.ContentID,
			&i.Key,
			&i.Type,
			&i.CreatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
