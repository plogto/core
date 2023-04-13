-- name: CreateLikedPost :one
INSERT INTO
	liked_posts (user_id, post_id)
VALUES
	($1, $2) RETURNING *;

-- name: GetLikedPostByUserIDAndPostID :one
SELECT
	*
FROM
	liked_posts
WHERE
	user_id = $1
	AND post_id = $2
	AND deleted_at IS NULL;

-- name: GetLikedPostByID :one
SELECT
	*
FROM
	liked_posts
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetLikedPostsByPostIDAndPageInfo :many
SELECT
	*
FROM
	liked_posts
WHERE
	post_id = $1
	AND created_at < $2
	AND deleted_at IS NULL
GROUP BY
	id,
	created_at
ORDER BY
	created_at DESC
LIMIT
	$3;

-- name: CountLikedPostsByPostIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		liked_posts
	WHERE
		post_id = $1
		AND created_at < $2
		AND deleted_at IS NULL
	GROUP BY
		created_at
	ORDER BY
		created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetLikedPostsByUserIDAndPageInfo :many
WITH _count_wrapper AS (
	SELECT
		liked_post.id AS liked_post_id,
		liked_post.post_id AS liked_post_post_id,
		liked_post.user_id AS liked_post_user_id,
		liked_post.created_at,
		liked_post.deleted_at AS liked_post_deleted_at,
		connections.id AS connection_id,
		connections.status,
		connections.deleted_at AS connection_deleted_at,
		posts.id AS post_id,
		posts.user_id AS post_user_id,
		posts.deleted_at AS post_deleted_at,
		users.id AS user_id,
		users.is_private
	FROM
		liked_posts AS liked_post
		INNER JOIN posts ON posts.id = liked_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		FULL OUTER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		liked_post.user_id = sqlc.arg(user_id)
		AND liked_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			connections.status = 2
			OR users.is_private = FALSE
		)
		AND (
			users.id = sqlc.arg(user_id)
			OR connections.deleted_at IS NULL
		)
		AND liked_post.created_at < sqlc.arg(created_at)
	GROUP BY
		connections.id,
		liked_post.id,
		posts.id,
		users.id
)
SELECT
	liked_post_id AS id,
	user_id,
	post_id,
	created_at
FROM
	_count_wrapper
GROUP BY
	id,
	user_id,
	post_id,
	created_at
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: CountLikedPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		liked_posts AS liked_post
		INNER JOIN posts ON posts.id = liked_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		FULL OUTER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		liked_post.user_id = sqlc.arg(user_id)
		AND liked_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			connections.status = 2
			OR users.is_private = FALSE
		)
		AND (
			users.id = sqlc.arg(user_id)
			OR connections.deleted_at IS NULL
		)
		AND liked_post.created_at < sqlc.arg(created_at)
	GROUP BY
		liked_post.id,
		posts.id,
		users.id
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: DeleteLikedPostByID :one
UPDATE
	liked_posts
SET
	deleted_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;