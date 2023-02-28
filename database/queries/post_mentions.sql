-- name: CreatePostMention :one
INSERT INTO
	post_mentions (user_id, post_id)
VALUES
	($1, $2) RETURNING *;

-- name: DeletePostMention :many
UPDATE
	post_mentions
SET
	deleted_at = $1
WHERE
	user_id = $2
	AND post_id = $3
	AND deleted_at IS NULL RETURNING *;

-- name: DeletePostMentionsByPostID :many
UPDATE
	post_mentions
SET
	deleted_at = $1
WHERE
	post_id = $2
	AND deleted_at IS NULL RETURNING *;