package model

import (
	"time"
)

type OnlineUser struct {
	tableName struct{}   `sql:"online_user"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Token     string     `json:"token"`
	SocketID  string     `json:"socket_id"`
	UserAgent string     `json:"user_agent"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:",soft_delete"`
}
