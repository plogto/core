-- name: CreateFile :one
INSERT INTO
	files (NAME, hash, width, height)
VALUES
	($1, $2, $3, $4) RETURNING *;

-- name: GetFileByID :one
SELECT
	*
FROM
	files
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetFileByName :one
SELECT
	*
FROM
	files
WHERE
	NAME = $1
	AND deleted_at IS NULL;

-- name: GetFileByHash :one
SELECT
	*
FROM
	files
WHERE
	hash = $1
	AND deleted_at IS NULL;

-- name: GetFilesByPostID :many
SELECT
	file.*
FROM
	files AS file
	INNER JOIN post_attachments ON post_attachments.file_id = file.id
WHERE
	post_attachments.post_id = $1
	AND file.deleted_at IS NULL
	AND post_attachments.deleted_at IS NULL
GROUP BY
	file.id;

-- name: GetFilesByTicketMessageID :many
SELECT
	file.*
FROM
	files AS file
	INNER JOIN ticket_message_attachments ON ticket_message_attachments.file_id = file.id
WHERE
	ticket_message_attachments.ticket_message_id = $1
	AND file.deleted_at IS NULL
	AND ticket_message_attachments.deleted_at IS NULL
GROUP BY
	file.id;