package model

import (
	"time"
)

type OnlineUser struct {
	tableName struct{} `pg:"online_user"`
	ID        string
	UserID    string
	Token     string
	SocketID  string
	UserAgent string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
