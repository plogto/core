package model

import (
	"time"
)

type File struct {
	tableName struct{} `pg:"file,discard_unknown_columns"`
	ID        string
	Hash      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
