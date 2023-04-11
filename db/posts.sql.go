// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: posts.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const countExplorePostsByPageInfo = `-- name: CountExplorePostsByPageInfo :one
WITH _count_wrapper AS (
	SELECT
		post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
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
	_count_wrapper
`

func (q *Queries) CountExplorePostsByPageInfo(ctx context.Context, createdAt time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, countExplorePostsByPageInfo, createdAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countExplorePostsWithAttachmentByPageInfo = `-- name: CountExplorePostsWithAttachmentByPageInfo :one
WITH _count_wrapper AS (
	SELECT
		post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
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
	_count_wrapper
`

func (q *Queries) CountExplorePostsWithAttachmentByPageInfo(ctx context.Context, createdAt time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, countExplorePostsWithAttachmentByPageInfo, createdAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByParentIDAndPageInfo = `-- name: CountPostsByParentIDAndPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.parent_id = $1
			AND post.deleted_at IS NULL
			AND post.created_at < $2
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			users.is_private = FALSE
			AND post.user_id = users.id
			AND post.parent_id = $1
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $2
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(*)
FROM
	_posts
`

type CountPostsByParentIDAndPageInfoParams struct {
	ParentID  uuid.NullUUID
	CreatedAt time.Time
}

func (q *Queries) CountPostsByParentIDAndPageInfo(ctx context.Context, arg CountPostsByParentIDAndPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsByParentIDAndPageInfo, arg.ParentID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByTagIDAndPageInfo = `-- name: CountPostsByTagIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		post.id, user_id, parent_id, child_id, status, content, url, post.created_at, post.updated_at, post.deleted_at, post_tags.id, tag_id, post_id, post_tags.created_at, post_tags.updated_at, post_tags.deleted_at, users.id, username, email, full_name, bio, role, is_private, avatar, background, primary_color, background_color, is_verified, invitation_code, users.created_at, users.updated_at, users.deleted_at
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
	_count_wrapper
`

type CountPostsByTagIDAndPageInfoParams struct {
	TagID     uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CountPostsByTagIDAndPageInfo(ctx context.Context, arg CountPostsByTagIDAndPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsByTagIDAndPageInfo, arg.TagID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByUserID = `-- name: CountPostsByUserID :one
SELECT
	count(*)
FROM
	posts
WHERE
	user_id = $1
	AND parent_id IS NULL
	AND deleted_at IS NULL
`

func (q *Queries) CountPostsByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsByUserID, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByUserIDAndPageInfo = `-- name: CountPostsByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
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
	_count_wrapper
`

type CountPostsByUserIDAndPageInfoParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CountPostsByUserIDAndPageInfo(ctx context.Context, arg CountPostsByUserIDAndPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsByUserIDAndPageInfo, arg.UserID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsByUserIDAndParentIDAndPageInfo = `-- name: CountPostsByUserIDAndParentIDAndPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.user_id = $1
			AND post.parent_id = $2
			AND post.deleted_at IS NULL
			AND post.created_at < $3
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = $1
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id = $2
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $3
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(*)
FROM
	_posts
`

type CountPostsByUserIDAndParentIDAndPageInfoParams struct {
	UserID    uuid.UUID
	ParentID  uuid.NullUUID
	CreatedAt time.Time
}

func (q *Queries) CountPostsByUserIDAndParentIDAndPageInfo(ctx context.Context, arg CountPostsByUserIDAndParentIDAndPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsByUserIDAndParentIDAndPageInfo, arg.UserID, arg.ParentID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countPostsWithParentIDByUserIDAndPageInfo = `-- name: CountPostsWithParentIDByUserIDAndPageInfo :one
WITH _count_wrapper AS (
	SELECT
		id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
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
	_count_wrapper
`

type CountPostsWithParentIDByUserIDAndPageInfoParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CountPostsWithParentIDByUserIDAndPageInfo(ctx context.Context, arg CountPostsWithParentIDByUserIDAndPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPostsWithParentIDByUserIDAndPageInfo, arg.UserID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countTimelinePostsByPageInfo = `-- name: CountTimelinePostsByPageInfo :one
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
		WHERE
			post.user_id = $1
			AND post.parent_id IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $2
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = $1
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
			AND post.created_at < $2
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	count(*)
FROM
	_posts
`

type CountTimelinePostsByPageInfoParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CountTimelinePostsByPageInfo(ctx context.Context, arg CountTimelinePostsByPageInfoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTimelinePostsByPageInfo, arg.UserID, arg.CreatedAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO
	posts (parent_id, user_id, CONTENT, status, url)
VALUES
	($1, $2, $3, $4, $5) RETURNING id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
`

type CreatePostParams struct {
	ParentID uuid.NullUUID
	UserID   uuid.UUID
	Content  sql.NullString
	Status   PostStatus
	Url      string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (*Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ParentID,
		arg.UserID,
		arg.Content,
		arg.Status,
		arg.Url,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentID,
		&i.ChildID,
		&i.Status,
		&i.Content,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deletePostByID = `-- name: DeletePostByID :one
UPDATE
	posts
SET
	deleted_at = $1
WHERE
	id = $2
	AND deleted_at IS NULL RETURNING id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
`

type DeletePostByIDParams struct {
	DeletedAt sql.NullTime
	ID        uuid.UUID
}

func (q *Queries) DeletePostByID(ctx context.Context, arg DeletePostByIDParams) (*Post, error) {
	row := q.db.QueryRowContext(ctx, deletePostByID, arg.DeletedAt, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentID,
		&i.ChildID,
		&i.Status,
		&i.Content,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getExplorePostsByPageInfo = `-- name: GetExplorePostsByPageInfo :many
SELECT
	post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
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
	$2
`

type GetExplorePostsByPageInfoParams struct {
	CreatedAt time.Time
	Limit     int32
}

func (q *Queries) GetExplorePostsByPageInfo(ctx context.Context, arg GetExplorePostsByPageInfoParams) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, getExplorePostsByPageInfo, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExplorePostsWithAttachmentByPageInfo = `-- name: GetExplorePostsWithAttachmentByPageInfo :many
SELECT
	post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
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
	$2
`

type GetExplorePostsWithAttachmentByPageInfoParams struct {
	CreatedAt time.Time
	Limit     int32
}

func (q *Queries) GetExplorePostsWithAttachmentByPageInfo(ctx context.Context, arg GetExplorePostsWithAttachmentByPageInfoParams) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, getExplorePostsWithAttachmentByPageInfo, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostByID = `-- name: GetPostByID :one
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	posts
WHERE
	id = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetPostByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentID,
		&i.ChildID,
		&i.Status,
		&i.Content,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getPostByURL = `-- name: GetPostByURL :one
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	posts
WHERE
	url = $1
	AND deleted_at IS NULL
`

func (q *Queries) GetPostByURL(ctx context.Context, url string) (*Post, error) {
	row := q.db.QueryRowContext(ctx, getPostByURL, url)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentID,
		&i.ChildID,
		&i.Status,
		&i.Content,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getPostsByIDs = `-- name: GetPostsByIDs :many
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	posts
WHERE
	id = ANY($1 :: uuid [ ])
	AND deleted_at IS NULL
`

func (q *Queries) GetPostsByIDs(ctx context.Context, dollar_1 []uuid.UUID) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByIDs, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByParentIDAndPageInfo = `-- name: GetPostsByParentIDAndPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.parent_id = $2
			AND post.deleted_at IS NULL
			AND post.created_at < $3
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			users.is_private = FALSE
			AND post.user_id = users.id
			AND post.parent_id = $2
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $3
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	_posts
LIMIT
	$1
`

type GetPostsByParentIDAndPageInfoParams struct {
	Limit     int32
	ParentID  uuid.NullUUID
	CreatedAt time.Time
}

type GetPostsByParentIDAndPageInfoRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ParentID  uuid.NullUUID
	ChildID   uuid.NullUUID
	Status    PostStatus
	Content   sql.NullString
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (q *Queries) GetPostsByParentIDAndPageInfo(ctx context.Context, arg GetPostsByParentIDAndPageInfoParams) ([]*GetPostsByParentIDAndPageInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByParentIDAndPageInfo, arg.Limit, arg.ParentID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetPostsByParentIDAndPageInfoRow{}
	for rows.Next() {
		var i GetPostsByParentIDAndPageInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByTagIDAndPageInfo = `-- name: GetPostsByTagIDAndPageInfo :many
SELECT
	post.id, user_id, parent_id, child_id, status, content, url, post.created_at, post.updated_at, post.deleted_at, post_tags.id, tag_id, post_id, post_tags.created_at, post_tags.updated_at, post_tags.deleted_at, users.id, username, email, full_name, bio, role, is_private, avatar, background, primary_color, background_color, is_verified, invitation_code, users.created_at, users.updated_at, users.deleted_at
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
	$3
`

type GetPostsByTagIDAndPageInfoParams struct {
	TagID     uuid.UUID
	CreatedAt time.Time
	Limit     int32
}

type GetPostsByTagIDAndPageInfoRow struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	ParentID        uuid.NullUUID
	ChildID         uuid.NullUUID
	Status          PostStatus
	Content         sql.NullString
	Url             string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
	ID_2            uuid.UUID
	TagID           uuid.UUID
	PostID          uuid.UUID
	CreatedAt_2     time.Time
	UpdatedAt_2     time.Time
	DeletedAt_2     sql.NullTime
	ID_3            uuid.UUID
	Username        string
	Email           string
	FullName        string
	Bio             sql.NullString
	Role            UserRole
	IsPrivate       bool
	Avatar          uuid.NullUUID
	Background      uuid.NullUUID
	PrimaryColor    PrimaryColor
	BackgroundColor BackgroundColor
	IsVerified      bool
	InvitationCode  string
	CreatedAt_3     time.Time
	UpdatedAt_3     time.Time
	DeletedAt_3     sql.NullTime
}

func (q *Queries) GetPostsByTagIDAndPageInfo(ctx context.Context, arg GetPostsByTagIDAndPageInfoParams) ([]*GetPostsByTagIDAndPageInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByTagIDAndPageInfo, arg.TagID, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetPostsByTagIDAndPageInfoRow{}
	for rows.Next() {
		var i GetPostsByTagIDAndPageInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.TagID,
			&i.PostID,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.DeletedAt_2,
			&i.ID_3,
			&i.Username,
			&i.Email,
			&i.FullName,
			&i.Bio,
			&i.Role,
			&i.IsPrivate,
			&i.Avatar,
			&i.Background,
			&i.PrimaryColor,
			&i.BackgroundColor,
			&i.IsVerified,
			&i.InvitationCode,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
			&i.DeletedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByUserIDAndPageInfo = `-- name: GetPostsByUserIDAndPageInfo :many
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
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
	$3
`

type GetPostsByUserIDAndPageInfoParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	Limit     int32
}

func (q *Queries) GetPostsByUserIDAndPageInfo(ctx context.Context, arg GetPostsByUserIDAndPageInfoParams) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUserIDAndPageInfo, arg.UserID, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsByUserIDAndParentIDAndPageInfo = `-- name: GetPostsByUserIDAndParentIDAndPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN users ON users.id = post.user_id
		WHERE
			post.user_id = $2
			AND post.parent_id = $3
			AND post.deleted_at IS NULL
			AND post.created_at < $4
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = $2
			INNER JOIN users ON users.id = connections.following_id
		WHERE
			(
				connections.status = 2
				OR users.is_private = FALSE
			)
			AND post.user_id = users.id
			AND post.parent_id = $3
			AND connections.deleted_at IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $4
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	_posts
LIMIT
	$1
`

type GetPostsByUserIDAndParentIDAndPageInfoParams struct {
	Limit     int32
	UserID    uuid.UUID
	ParentID  uuid.NullUUID
	CreatedAt time.Time
}

type GetPostsByUserIDAndParentIDAndPageInfoRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ParentID  uuid.NullUUID
	ChildID   uuid.NullUUID
	Status    PostStatus
	Content   sql.NullString
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (q *Queries) GetPostsByUserIDAndParentIDAndPageInfo(ctx context.Context, arg GetPostsByUserIDAndParentIDAndPageInfoParams) ([]*GetPostsByUserIDAndParentIDAndPageInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostsByUserIDAndParentIDAndPageInfo,
		arg.Limit,
		arg.UserID,
		arg.ParentID,
		arg.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetPostsByUserIDAndParentIDAndPageInfoRow{}
	for rows.Next() {
		var i GetPostsByUserIDAndParentIDAndPageInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPostsWithParentIDByUserIDAndPageInfo = `-- name: GetPostsWithParentIDByUserIDAndPageInfo :many
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
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
	$3
`

type GetPostsWithParentIDByUserIDAndPageInfoParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	Limit     int32
}

func (q *Queries) GetPostsWithParentIDByUserIDAndPageInfo(ctx context.Context, arg GetPostsWithParentIDByUserIDAndPageInfoParams) ([]*Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsWithParentIDByUserIDAndPageInfo, arg.UserID, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTimelinePostsByPageInfo = `-- name: GetTimelinePostsByPageInfo :many
WITH _posts AS (
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
		WHERE
			post.user_id = $2
			AND post.parent_id IS NULL
			AND post.deleted_at IS NULL
			AND post.created_at < $3
	)
	UNION
	(
		SELECT
			post.id, post.user_id, post.parent_id, post.child_id, post.status, post.content, post.url, post.created_at, post.updated_at, post.deleted_at
		FROM
			posts AS post
			INNER JOIN connections ON connections.follower_id = $2
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
			AND post.created_at < $3
		ORDER BY
			post.created_at DESC
	)
)
SELECT
	id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
FROM
	_posts
LIMIT
	$1
`

type GetTimelinePostsByPageInfoParams struct {
	Limit     int32
	UserID    uuid.UUID
	CreatedAt time.Time
}

type GetTimelinePostsByPageInfoRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ParentID  uuid.NullUUID
	ChildID   uuid.NullUUID
	Status    PostStatus
	Content   sql.NullString
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (q *Queries) GetTimelinePostsByPageInfo(ctx context.Context, arg GetTimelinePostsByPageInfoParams) ([]*GetTimelinePostsByPageInfoRow, error) {
	rows, err := q.db.QueryContext(ctx, getTimelinePostsByPageInfo, arg.Limit, arg.UserID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetTimelinePostsByPageInfoRow{}
	for rows.Next() {
		var i GetTimelinePostsByPageInfoRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ParentID,
			&i.ChildID,
			&i.Status,
			&i.Content,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE
	posts
SET
	CONTENT = $1,
	status = $2
WHERE
	id = $3
	AND deleted_at IS NULL RETURNING id, user_id, parent_id, child_id, status, content, url, created_at, updated_at, deleted_at
`

type UpdatePostParams struct {
	Content sql.NullString
	Status  PostStatus
	ID      uuid.UUID
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (*Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.Content, arg.Status, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentID,
		&i.ChildID,
		&i.Status,
		&i.Content,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
