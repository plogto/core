package model

import "time"

type Ticket struct {
	tableName struct{} `pg:"tickets"`
	ID        string
	UserID    string
	Subject   string
	Status    TicketStatus
	Url       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
