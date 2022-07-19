package model

import (
	"time"
)

type File struct {
	tableName struct{} `pg:"files,discard_unknown_columns"`
	ID        string
	Hash      string
	Name      string
	Width     int32
	Height    int32
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
