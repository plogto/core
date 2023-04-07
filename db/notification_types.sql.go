// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: notification_types.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getNotificationTypeByID = `-- name: GetNotificationTypeByID :one
SELECT
	id, name, template, deleted_at
FROM
	notification_types
WHERE
	id = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetNotificationTypeByID(ctx context.Context, id uuid.UUID) (*NotificationType, error) {
	row := q.db.QueryRowContext(ctx, getNotificationTypeByID, id)
	var i NotificationType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Template,
		&i.DeletedAt,
	)
	return &i, err
}

const getNotificationTypeByName = `-- name: GetNotificationTypeByName :one
SELECT
	id, name, template, deleted_at
FROM
	notification_types
WHERE
	NAME = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetNotificationTypeByName(ctx context.Context, name NotificationTypeName) (*NotificationType, error) {
	row := q.db.QueryRowContext(ctx, getNotificationTypeByName, name)
	var i NotificationType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Template,
		&i.DeletedAt,
	)
	return &i, err
}