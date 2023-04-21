-- name: CreateUser :one
INSERT INTO
	users (
		email,
		username,
		invitation_code,
		full_name,
		settings
	)
VALUES
	($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUsersByIDs :many
SELECT
	*
FROM
	users
WHERE
	id = ANY($1 :: uuid [ ])
	AND deleted_at IS NULL;

-- name: GetUserByID :one
SELECT
	*
FROM
	users
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetUserByUsername :one
SELECT
	*
FROM
	users
WHERE
	lower(username) = lower($1)
	AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT
	*
FROM
	users
WHERE
	lower(email) = lower($1)
	AND deleted_at IS NULL;

-- name: GetUserByInvitationCode :one
SELECT
	*
FROM
	users
WHERE
	lower(invitation_code) = lower($1)
	AND deleted_at IS NULL;

-- name: GetUserByUsernameOrEmail :one
SELECT
	*
FROM
	users
WHERE
	(
		lower(email) = lower($1)
		OR lower(username) = lower($1)
	)
	AND deleted_at IS NULL;

-- name: GetUsersByUsernameOrFullNameAndPageInfo :many
SELECT
	*
FROM
	users
WHERE
	(
		lower(email) LIKE lower($1)
		OR lower(username) LIKE lower($1)
	)
	AND deleted_at IS NULL
LIMIT
	$2;

-- name: UpdateUser :one
UPDATE
	users
SET
	username = sqlc.arg(username),
	email = sqlc.arg(email),
	bio = sqlc.arg(bio),
	full_name = sqlc.arg(full_name),
	background = sqlc.arg(background),
	avatar = sqlc.arg(avatar),
	background_color = sqlc.arg(background_color),
	primary_color = sqlc.arg(primary_color),
	is_private = sqlc.arg(is_private)
WHERE
	id = $1
	AND deleted_at IS NULL RETURNING *;

-- name: UpdateUserSettings :one
UPDATE
	users
SET
	settings = sqlc.arg(settings)
WHERE
	id = $1
	AND deleted_at IS NULL RETURNING *;