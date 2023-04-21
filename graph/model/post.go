package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

type Post struct {
	ID         pgtype.UUID
	Status     PostStatus
	Parent     *Post
	ParentID   pgtype.UUID
	Child      *Post
	ChildID    pgtype.UUID
	User       db.User
	UserID     pgtype.UUID
	Content    pgtype.Text
	Attachment []*db.File
	Url        string
	Likes      *LikedPosts
	Replies    *Posts
	IsLiked    *db.LikedPost
	IsSaved    *db.SavedPost
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
