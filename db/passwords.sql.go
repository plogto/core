// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: passwords.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPassword = `-- name: CreatePassword :one
INSERT INTO
	passwords (user_id, PASSWORD)
VALUES
	($1, $2) RETURNING id, user_id, password, created_at, updated_at, deleted_at
`

type CreatePasswordParams struct {
	UserID   pgtype.UUID
	Password string
}

func (q *Queries) CreatePassword(ctx context.Context, arg CreatePasswordParams) (*Password, error) {
	row := q.db.QueryRow(ctx, createPassword, arg.UserID, arg.Password)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getPasswordByUserID = `-- name: GetPasswordByUserID :one
SELECT
	id, user_id, password, created_at, updated_at, deleted_at
FROM
	passwords
WHERE
	user_id = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetPasswordByUserID(ctx context.Context, userID pgtype.UUID) (*Password, error) {
	row := q.db.QueryRow(ctx, getPasswordByUserID, userID)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE
	passwords
SET
	PASSWORD = $1
WHERE
	user_id = $2
	AND deleted_at IS NULL RETURNING id, user_id, password, created_at, updated_at, deleted_at
`

type UpdatePasswordParams struct {
	Password string
	UserID   pgtype.UUID
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (*Password, error) {
	row := q.db.QueryRow(ctx, updatePassword, arg.Password, arg.UserID)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
