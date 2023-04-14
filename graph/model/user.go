package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type Post struct {
	ID         uuid.UUID
	Status     PostStatus
	Parent     *Post
	ParentID   uuid.NullUUID
	Child      *Post
	ChildID    uuid.NullUUID
	User       db.User
	UserID     uuid.UUID
	Content    sql.NullString
	Attachment []*db.File
	Url        string
	Likes      *LikedPosts
	Replies    *Posts
	IsLiked    *db.LikedPost
	IsSaved    *db.SavedPost
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
