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

-- name: GetSavedPostsByPostIDAndPageInfo :many
SELECT
	*
FROM
	saved_posts
WHERE
	post_id = $1
	AND created_at < $2
	AND deleted_at IS NULL
ORDER BY
	created_at DESC
LIMIT
	$3;

-- name: CountSavedPostsByPostIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		saved_posts
	WHERE
		post_id = $1
		AND created_at < $2
		AND deleted_at IS NULL
	ORDER BY
		created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetSavedPostsByUserIDAndPageInfo :many
SELECT
	*
FROM
	saved_posts AS saved_post
	INNER JOIN posts ON posts.id = saved_post.post_id
	INNER JOIN users ON users.id = posts.user_id
	INNER JOIN connections ON connections.following_id = posts.user_id
WHERE
	saved_post.user_id = $1
	AND saved_post.deleted_at IS NULL
	AND posts.deleted_at IS NULL
	AND (
		users.id = $1
		OR connections.status = 2
		OR users.is_private = FALSE
	)
	AND connections.deleted_at IS NULL
	AND saved_post.created_at < $2
GROUP BY
	connections.id,
	saved_post.id,
	posts.id,
	users.id
LIMIT
	$3;

-- name: CountSavedPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		saved_posts AS saved_post
		INNER JOIN posts ON posts.id = saved_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		INNER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		saved_post.user_id = $1
		AND saved_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			users.id = $1
			OR connections.status = 2
			OR users.is_private = FALSE
		)
		AND connections.deleted_at IS NULL
		AND saved_post.created_at < $2
	GROUP BY
		connections.id,
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