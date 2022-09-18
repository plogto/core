package model

import (
	"time"
)

type TicketMessageAttachment struct {
	tableName       struct{} `pg:"ticket_message_attachments"`
	ID              string
	TicketMessageID string
	FileID          string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time `pg:"-,soft_delete"`
}
