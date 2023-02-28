-- name: CreateNotification :one
INSERT INTO
	notifications (
		notification_type_id,
		sender_id,
		receiver_id,
		deleted_at,
		post_id,
		reply_id,
		url
	)
VALUES
	($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetNotification :one
SELECT
	*
FROM
	notifications
WHERE
	notification_type_id = $1
	AND sender_id = $2
	AND receiver_id = $3
	AND post_id = $4
	AND reply_id = $5
	AND url = $6
	AND deleted_at IS NULL
LIMIT
	1;

-- name: GetNotificationByID :one
SELECT
	*
FROM
	notifications
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetNotificationsByReceiverIDAndPageInfo :many
SELECT
	*
FROM
	notifications
WHERE
	receiver_id = $1
	AND deleted_at IS NULL
	AND created_at < $2
LIMIT
	$3;

-- name: CountNotificationsByReceiverIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		notifications
	WHERE
		receiver_id = $1
		AND deleted_at IS NULL
		AND created_at < $2
	LIMIT
		$3
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: CountUnreadNotificationsByReceiverID :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		notifications
	WHERE
		receiver_id = $1
		AND READ = FALSE
		AND deleted_at IS NULL
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: UpdateReadNotifications :many
UPDATE
	notifications
SET
	READ = TRUE
WHERE
	receiver_id = $1
	AND deleted_at IS NULL RETURNING *;

-- name: RemovePostNotificationsByPostID :many
UPDATE
	notifications
SET
	deleted_at = $1
WHERE
	post_id = $2
	AND deleted_at IS NULL RETURNING *;

-- name: RemoveNotification :one
UPDATE
	notifications
SET
	deleted_at = $1
WHERE
	notification_type_id = $2
	AND sender_id = $3
	AND receiver_id = $4
	AND post_id = $5
	AND reply_id = $6
	AND deleted_at IS NULL RETURNING *;