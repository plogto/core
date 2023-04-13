-- name: CreateSavedPost :one
INSERT INTO
	saved_posts (user_id, post_id)
VALUES
	($1, $2) RETURNING *;

-- name: GetSavedPostByUserIDAndPostID :one
SELECT
	*
FROM
	saved_posts
WHERE
	user_id = $1
	AND post_id = $2
	AND deleted_at IS NULL;

-- name: GetSavedPostByID :one
SELECT
	*
FROM
	saved_posts
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetSavedPostsByUserIDAndPageInfo :many
WITH _count_wrapper AS (
	SELECT
		saved_post.id AS saved_post_id,
		saved_post.post_id AS saved_post_post_id,
		saved_post.user_id AS saved_post_user_id,
		saved_post.created_at,
		saved_post.deleted_at AS saved_post_deleted_at,
		connections.id AS connection_id,
		connections.status,
		connections.deleted_at AS connection_deleted_at,
		posts.id AS post_id,
		posts.user_id AS post_user_id,
		posts.deleted_at AS post_deleted_at,
		users.id AS user_id,
		users.is_private
	FROM
		saved_posts AS saved_post
		INNER JOIN posts ON posts.id = saved_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		FULL OUTER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		saved_post.user_id = sqlc.arg(user_id)
		AND saved_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			connections.status = 2
			OR users.is_private = FALSE
		)
		AND (
			users.id = sqlc.arg(user_id)
			OR connections.deleted_at IS NULL
		)
		AND saved_post.created_at < sqlc.arg(created_at)
	GROUP BY
		connections.id,
		saved_post.id,
		posts.id,
		users.id
)
SELECT
	saved_post_id AS id,
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

-- name: CountSavedPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		saved_posts AS saved_post
		INNER JOIN posts ON posts.id = saved_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		FULL OUTER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		saved_post.user_id = sqlc.arg(user_id)
		AND saved_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			connections.status = 2
			OR users.is_private = FALSE
		)
		AND (
			users.id = sqlc.arg(user_id)
			OR connections.deleted_at IS NULL
		)
		AND saved_post.created_at < sqlc.arg(created_at)
	GROUP BY
		saved_post.id,
		posts.id,
		users.id
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: DeleteSavedPostByID :one
UPDATE
	saved_posts
SET
	deleted_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;