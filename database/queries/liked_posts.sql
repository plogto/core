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
SELECT
	*
FROM
	liked_posts AS liked_post
	INNER JOIN posts ON posts.id = liked_post.post_id
	INNER JOIN users ON users.id = posts.user_id
	INNER JOIN connections ON connections.following_id = posts.user_id
WHERE
	(liked_post.user_id = $1)
	AND (liked_post.deleted_at IS NULL)
	AND (posts.deleted_at IS NULL)
	AND (
		(users.id = $1)
		OR (connections.status = 2)
		OR (users.is_private = FALSE)
	)
	AND (connections.deleted_at IS NULL)
	AND liked_post.created_at < $2
GROUP BY
	connections.id,
	liked_post.id,
	posts.id,
	users.id
LIMIT
	$3;

-- name: CountLikedPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		count(*)
	FROM
		liked_posts AS liked_post
		INNER JOIN posts ON posts.id = liked_post.post_id
		INNER JOIN users ON users.id = posts.user_id
		INNER JOIN connections ON connections.following_id = posts.user_id
	WHERE
		liked_post.user_id = $1
		AND liked_post.deleted_at IS NULL
		AND posts.deleted_at IS NULL
		AND (
			users.id = $1
			OR connections.status = 2
			OR users.is_private = FALSE
		)
		AND connections.deleted_at IS NULL
		AND liked_post.created_at < $2
	GROUP BY
		connections.id,
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