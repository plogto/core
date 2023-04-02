-- name: CreateConnection :one
INSERT INTO
	connections (following_id, follower_id, status)
VALUES
	($1, $2, $3) RETURNING *;

-- name: GetConnectionByID :one
SELECT
	*
FROM
	connections
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetConnection :one
SELECT
	*
FROM
	connections
WHERE
	following_id = $1
	AND follower_id = $2
	AND deleted_at IS NULL;

-- name: GetConnectionsByIDs :many
SELECT
	*
FROM
	connections
WHERE
	id = ANY($1 :: uuid [ ])
	AND deleted_at IS NULL;

-- name: GetFollowersByUserIDAndPageInfo :many
SELECT
	*
FROM
	connections
WHERE
	following_id = $1
	AND status = $2
	AND created_at < $3
	AND deleted_at IS NULL
GROUP BY
	id
ORDER BY
	created_at DESC
LIMIT
	$4;

-- name: CountFollowersByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		connections
	WHERE
		following_id = $1
		AND status = $2
		AND created_at < $3
		AND deleted_at IS NULL
	GROUP BY
		id
	ORDER BY
		created_at DESC
	LIMIT
		$4
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetFollowingByUserIDAndPageInfo :many
SELECT
	*
FROM
	connections
WHERE
	follower_id = $1
	AND status = $2
	AND created_at < $3
	AND deleted_at IS NULL
GROUP BY
	id
ORDER BY
	created_at DESC
LIMIT
	$4;

-- name: CountFollowingByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		connections
	WHERE
		follower_id = $1
		AND status = $2
		AND created_at < $3
		AND deleted_at IS NULL
	GROUP BY
		id
	ORDER BY
		created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: UpdateConnection :one
UPDATE
	connections
SET
	status = $1
WHERE
	follower_id = $2
	AND following_id = $3
	AND deleted_at IS NULL RETURNING *;

-- name: DeleteConnection :one
UPDATE
	connections
SET
	deleted_at = $1
WHERE
	id = $2 RETURNING *;