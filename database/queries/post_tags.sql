-- name: CreatePostTag :one
INSERT INTO
	post_tags (tag_id, post_id)
VALUES
	($1, $2) RETURNING *;

-- name: CountPostTagsByTagID :one
WITH _count_wrapper AS (
	SELECT
		post_tag.*
	FROM
		post_tags AS post_tag
		INNER JOIN posts ON posts.id = post_tag.post_id
		INNER JOIN users ON users.id = posts.user_id
	WHERE
		post_tag.tag_id = $1
		AND posts.deleted_at IS NULL
		AND users.is_private IS FALSE
	GROUP BY
		post_tag.tag_id,
		post_tag.id
)
SELECT
	count(*)
FROM
	_count_wrapper;

-- name: GetTagsOrderByCountTags :many
SELECT
	tag.*,
	count(tag.id)
FROM
	tags AS tag
	INNER JOIN post_tags ON post_tags.tag_id = tag.id
	INNER JOIN posts ON post_tags.post_id = posts.id
	INNER JOIN users ON users.id = posts.user_id
WHERE
	posts.deleted_at IS NULL
	AND users.is_private IS FALSE
GROUP BY
	tag.id,
	post_tags.tag_id
ORDER BY
	count DESC,
	tag.created_at DESC
LIMIT
	$1;

-- name: DeletePostTagsByPostID :many
UPDATE
	post_tags
SET
	deleted_at = $1
WHERE
	post_id = $2
	AND deleted_at IS NULL RETURNING *;