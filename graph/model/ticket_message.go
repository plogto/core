package model

import "time"

type TicketMessage struct {
	tableName struct{} `pg:"ticket_messages"`
	ID        string
	TicketID  string
	SenderID  string
	Message   string
	Read      *bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
