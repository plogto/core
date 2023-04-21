package model

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Tag struct {
	ID        pgtype.UUID
	Name      string
	Count     int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
