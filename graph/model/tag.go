package model

import (
	"time"
)

type Tag struct {
	tableName struct{} `pg:"tags,discard_unknown_columns"`
	ID        string
	Name      string
	Count     int64 `pg:"-"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
