-- name: CreatePostAttachment :one
INSERT INTO
	post_attachments (post_id, file_id)
VALUES
	($1, $2) RETURNING *;