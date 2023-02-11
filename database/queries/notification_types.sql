-- name: GetNotificationTypeByID :one
SELECT
	*
FROM
	notification_types
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetNotificationTypeByName :one
SELECT
	*
FROM
	notification_types
WHERE
	NAME = $1
	AND deleted_at IS NULL;