package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID
	Name      string
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
