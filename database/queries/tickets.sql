-- name: CreateTicket :one
INSERT INTO
	tickets (user_id, subject, url)
VALUES
	($1, $2, $3) RETURNING *;

-- name: GetTicketByID :one
SELECT
	*
FROM
	tickets
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetTicketByURL :one
SELECT
	*
FROM
	tickets
WHERE
	url = $1
	AND deleted_at IS NULL;

-- name: GetTicketsByUserIDAndPageInfo :many
SELECT
	*
FROM
	tickets
WHERE
	(
		user_id = sqlc.narg(user_id)
		OR sqlc.narg(user_id) IS NULL
	)
	AND updated_at < $1
	AND deleted_at IS NULL
ORDER BY
	updated_at DESC
LIMIT
	$2;

-- name: CountTicketsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		*
	FROM
		tickets
	WHERE
		(
			user_id = sqlc.narg(user_id)
			OR sqlc.narg(user_id) IS NULL
		)
		AND updated_at < $1
		AND deleted_at IS NULL
	ORDER BY
		updated_at DESC
	LIMIT
		$2
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: UpdateTicketStatus :one
UPDATE
	tickets
SET
	status = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;

-- name: UpdateTicketUpdatedAt :one
UPDATE
	tickets
SET
	updated_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;