-- name: CreatePost :one
INSERT INTO
	posts (parent_id, child_id, user_id, CONTENT, status, url)
VALUES
	($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetPostsByIDs :many
SELECT
	*
FROM
	posts
WHERE
	id = ANY($1 :: uuid [ ])
	AND deleted_at IS NULL;

-- name: GetChildPostsByIDsAndUserID :many
SELECT
	*
FROM
	posts
WHERE
	user_id = sqlc.arg(user_id)
	AND child_id = ANY(sqlc.arg(child_id) :: uuid [ ])
	AND deleted_at IS NULL;

-- name: GetPostByID :one
SELECT
	*
FROM
	posts
WHERE
	id = $1
	AND deleted_at IS NULL;

-- name: GetPostByURL :one
SELECT
	*
FROM
	posts
WHERE
	url = $1
	AND deleted_at IS NULL;

-- name: GetPostsByUserIDAndPageInfo :many
SELECT
	*
FROM
	posts
WHERE
	user_id = $1
	AND deleted_at IS NULL
	AND parent_id IS NULL
	AND created_at < $2
ORDER BY
	created_at DESC
LIMIT
	$3;

-- name: CountPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		*
	FROM
		posts
	WHERE
		user_id = $1
		AND deleted_at IS NULL
		AND parent_id IS NULL
		AND created_at < $2
	ORDER BY
		created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetPostsWithAttachmentByUserIDAndPageInfo :many
SELECT
	DISTINCT ON (post.created_at) post.id,
	post.*,
	post_attachments.*
FROM
	posts AS post
	INNER JOIN post_attachments ON post_attachments.post_id = post.id
WHERE
	post.user_id = sqlc.arg(user_id)
	AND post.deleted_at IS NULL
	AND post.created_at < sqlc.arg(created_at)
GROUP BY
	post_attachments.id,
	post.id
ORDER BY
	post.created_at DESC
LIMIT
	$1;

-- name: CountPostsWithAttachmentByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		DISTINCT ON (post.created_at) post.id,
		post.*,
		post_attachments.*
	FROM
		posts AS post
		INNER JOIN post_attachments ON post_attachments.post_id = post.id
	WHERE
		post.user_id = sqlc.arg(user_id)
		AND post.deleted_at IS NULL
		AND post.created_at < sqlc.arg(created_at)
	GROUP BY
		post_attachments.id,
		post.id
	ORDER BY
		post.created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetPostsWithParentIDByUserIDAndPageInfo :many
SELECT
	*
FROM
	posts
WHERE
	user_id = $1
	AND deleted_at IS NULL
	AND parent_id IS NOT NULL
	AND created_at < $2
ORDER BY
	created_at DESC
LIMIT
	$3;

-- name: CountPostsWithParentIDByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		*
	FROM
		posts
	WHERE
		user_id = $1
		AND deleted_at IS NULL
		AND parent_id IS NOT NULL
		AND created_at < $2
	ORDER BY
		created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetPostsByTagIDAndPageInfo :many
SELECT
	*
FROM
	posts AS post
	INNER JOIN post_tags ON post_tags.tag_id = $1
	INNER JOIN users ON users.id = post.user_id
WHERE
	post_tags.post_id = post.id
	AND post.deleted_at IS NULL
	AND users.is_private IS FALSE
	AND post.created_at < $2
ORDER BY
	post.created_at DESC
LIMIT
	$3;

-- name: CountPostsByTagIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		*
	FROM
		posts AS post
		INNER JOIN post_tags ON post_tags.tag_id = $1
		INNER JOIN users ON users.id = post.user_id
	WHERE
		post_tags.post_id = post.id
		AND post.deleted_at IS NULL
		AND users.is_private IS FALSE
		AND post.created_at < $2
	ORDER BY
		post.created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetExplorePostsByPageInfo :many
SELECT
	post.*
FROM
	posts AS post
	INNER JOIN users ON users.id = post.user_id
	INNER JOIN post_tags ON post_tags.post_id = post.id
WHERE
	post.parent_id IS NULL
	AND users.is_private = FALSE
	AND post.deleted_at IS NULL
	AND post_tags.id IS NOT NULL
	AND post.created_at < $1
GROUP BY
	post.id
ORDER BY
	post.created_at DESC
LIMIT
	$2;

-- name: CountExplorePostsByPageInfo :one
WITH _count_wrapper AS (
	SELECT
		post.*
	FROM
		posts AS post
		INNER JOIN users ON users.id = post.user_id
		INNER JOIN post_tags ON post_tags.post_id = post.id
	WHERE
		post.parent_id IS NULL
		AND users.is_private = FALSE
		AND post.deleted_at IS NULL
		AND post_tags.id IS NOT NULL
		AND post.created_at < $1
	GROUP BY
		post.id
	ORDER BY
		post.created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetExplorePostsWithAttachmentByPageInfo :many
SELECT
	post.*
FROM
	posts AS post
	INNER JOIN users ON users.id = post.user_id
	INNER JOIN post_tags ON post_tags.post_id = post.id
	INNER JOIN post_attachments ON post_attachments.post_id = post.id
WHERE
	post.parent_id IS NULL
	AND users.is_private = FALSE
	AND post.deleted_at IS NULL
	AND post_tags.id IS NOT NULL
	AND post.created_at < $1
GROUP BY
	post.id
ORDER BY
	post.created_at DESC
LIMIT
	$2;

-- name: CountExplorePostsWithAttachmentByPageInfo :one
WITH _count_wrapper AS (
	SELECT
		post.*
	FROM
		posts AS post
		INNER JOIN users ON users.id = post.user_id
		INNER JOIN post_tags ON post_tags.post_id = post.id
		INNER JOIN post_attachments ON post_attachments.post_id = post.id
	WHERE
		post.parent_id IS NULL
		AND users.is_private = FALSE
		AND post.deleted_at IS NULL
		AND post_tags.id IS NOT NULL
		AND post.created_at < $1
	GROUP BY
		post.id
	ORDER BY
		post.created_at DESC
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: CountPostsByUserID :one
SELECT
	count(*)
FROM
	posts
WHERE
	user_id = $1
	AND parent_id IS NULL
	AND deleted_at IS NULL;

-- name: GetPostsByUserIDAndParentIDAndPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.user_id = sqlc.arg(user_id)
			AND post.parent_id = sqlc.narg(parent_id)
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = sqlc.arg(user_id)
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id = sqlc.narg(parent_id)
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	*
FROM
	_posts
LIMIT
	$1;

-- name: CountPostsByUserIDAndParentIDAndPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.user_id = sqlc.arg(user_id)
			AND post.parent_id = sqlc.narg(parent_id)
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = sqlc.arg(user_id)
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id = sqlc.narg(parent_id)
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(*)
FROM
	_posts;

-- name: GetPostsByParentIDAndPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.parent_id = sqlc.narg(parent_id)
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			users.is_private = FALSE
			AND post.user_id = users.id
			AND post.parent_id = sqlc.narg(parent_id)
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	*
FROM
	_posts
LIMIT
	$1;

-- name: CountPostsByParentIDAndPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.parent_id = sqlc.narg(parent_id)
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			users.is_private = FALSE
			AND post.user_id = users.id
			AND post.parent_id = sqlc.narg(parent_id)
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(*)
FROM
	_posts;

-- name: GetTimelinePostsByPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
		WHERE
			post.user_id = sqlc.arg(user_id)
			AND post.parent_id IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		GROUP BY
			post.id,
			post.created_at
		ORDER BY
			post.created_at DESC
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = sqlc.arg(user_id)
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id IS NULL
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		GROUP BY
			post.id,
			post.created_at
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	*
FROM
	_posts
ORDER BY
	created_at DESC
LIMIT
	$1;

-- name: CountTimelinePostsByPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.*
		FROM
			posts AS post
		WHERE
			post.user_id = sqlc.arg(user_id)
			AND post.parent_id IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		GROUP BY
			post.id,
			post.created_at
		ORDER BY
			post.created_at DESC
	)
	UNION
	(
		SELECT
			post.*
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = sqlc.arg(user_id)
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id IS NULL
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < sqlc.arg(created_at)
		GROUP BY
			post.id,
			post.created_at
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(id)
FROM
	_posts;

-- name: UpdatePost :one
UPDATE
	posts
SET
	CONTENT = $1,
	status = $2
WHERE
	id = $3
	AND deleted_at IS NULL RETURNING *;

-- name: DeletePostByID :one
UPDATE
	posts
SET
	deleted_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING *;