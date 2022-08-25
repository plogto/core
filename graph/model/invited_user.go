package model

import (
	"time"
)

type InvitedUser struct {
	tableName struct{} `pg:"invited_users"`
	ID        string
	InviterID string
	InviteeID string
	Status    *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
