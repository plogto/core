package model

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

type AddPostInputType struct {
	ParentID   uuid.UUID
	Content    sql.NullString
	Status     db.PostStatus
	Attachment []string
}
