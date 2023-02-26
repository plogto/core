-- name: GetPasswordByUserID :one
SELECT
	*
FROM
	passwords
WHERE
	user_id = $1
	AND deleted_at IS NULL;

-- name: CreatePassword :one
INSERT INTO
	passwords (user_id, PASSWORD)
VALUES
	($1, $2) RETURNING *;

-- name: UpdatePassword :one
UPDATE
	passwords
SET
	PASSWORD = $1
WHERE
	user_id = $2
	AND deleted_at IS NULL RETURNING *;