-- name: CreateTicketMessage :one
INSERT INTO
	ticket_messages (sender_id, ticket_id, message)
VALUES
	($1, $2, $3) RETURNING *;

-- name: GetTicketMessageByID :one
SELECT
	*
FROM
	ticket_messages
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetLastTicketMessageByTicketID :one
SELECT
	*
FROM
	ticket_messages
WHERE
	ticket_id = $1
	AND deleted_at IS NULL
ORDER BY
	created_at DESC
LIMIT
	1;

-- name: GetTicketMessagesByTicketIDAndPageInfo :many
SELECT
	*
FROM
	ticket_messages
WHERE
	ticket_id = $1
	AND created_at < $2
	AND deleted_at IS NULL
ORDER BY
	created_at DESC
LIMIT
	$3;

-- name: CountTicketMessagesByTicketIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		*
	FROM
		ticket_messages
	WHERE
		ticket_id = $1
		AND created_at < $2
		AND deleted_at IS NULL
	ORDER BY
		created_at DESC
	LIMIT
		$3
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: UpdateReadTicketMessagesByUserIDAndTicketID :many
UPDATE
	ticket_messages
SET
	READ = TRUE
WHERE
	sender_id = $1
	AND ticket_id = $2
	AND deleted_at IS NULL RETURNING id;