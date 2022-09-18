package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type TicketMessageAttachments struct {
	DB *pg.DB
}

func (t *TicketMessageAttachments) CreateTicketMessageAttachment(ticketMessageAttachment *model.TicketMessageAttachment) (*model.TicketMessageAttachment, error) {
	_, err := t.DB.Model(ticketMessageAttachment).Returning("*").Insert()
	return ticketMessageAttachment, err
}
