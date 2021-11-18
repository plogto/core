package model

import (
	"time"
)

type Tag struct {
	tableName struct{} `pg:"tag"`
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
