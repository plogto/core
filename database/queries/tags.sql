-- name: CreateTag :one
INSERT INTO
	tags (NAME)
VALUES
	($1) RETURNING *;

-- name: GetTagByIDs :many
SELECT
	*
FROM
	tags
WHERE
	id = ANY(sqlc.arg(ids) :: uuid [ ])
	AND deleted_at IS NULL;

-- name: GetTagByName :one
SELECT
	*
FROM
	tags
WHERE
	lower(tag.name) = lower(sqlc.arg(NAME))
	AND deleted_at IS NULL;

-- name: GetTagsByTagNameAndPageInfo :many
SELECT
	tag.*,
	count(tag.id),
	post_tags.tag_id
FROM
	tags AS tag
	INNER JOIN post_tags ON post_tags.tag_id = tag.id
	INNER JOIN posts ON post_tags.post_id = posts.id
	INNER JOIN users ON users.id = posts.user_id
WHERE
	lower(tag.name) LIKE lower(sqlc.arg(NAME))
	AND posts.deleted_at IS NULL
	AND users.is_private IS FALSE
GROUP BY
	tag.id,
	post_tags.tag_id
ORDER BY
	count DESC,
	tag.created_at DESC
LIMIT
	$1;