package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type TicketMessages struct {
	DB *pg.DB
}

func (t *TicketMessages) GetTicketMessageByField(field string, value string) (*model.TicketMessage, error) {
	var ticketMessage model.TicketMessage
	err := t.DB.Model(&ticketMessage).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(ticketMessage.ID) < 1 {
		return nil, nil
	}
	return &ticketMessage, err
}

func (t *TicketMessages) GetTicketMessagesByTicketIDAndPageInfo(ticketID string, limit int, after string) (*model.TicketMessages, error) {
	var ticketMessages []*model.TicketMessage
	var edges []*model.TicketMessagesEdge
	var endCursor string

	query := t.DB.Model(&ticketMessages).
		Where("ticket_id = ?", ticketID).
		Where("deleted_at is ?", nil).
		Order("created_at DESC")

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range ticketMessages {
		edges = append(edges, &model.TicketMessagesEdge{Node: &model.TicketMessage{
			ID:        value.ID,
			TicketID:  value.TicketID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.TicketMessages{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (t *TicketMessages) GetTicketMessageByID(id string) (*model.TicketMessage, error) {
	return t.GetTicketMessageByField("id", id)
}

func (t *TicketMessages) GetLastTicketMessageByTicketID(ticketID string) (*model.TicketMessage, error) {
	var ticketMessage model.TicketMessage
	err := t.DB.Model(&ticketMessage).Where("ticket_id = ?", ticketID).Where("deleted_at is ?", nil).Order("created_at DESC").First()
	if len(ticketMessage.ID) < 1 {
		return nil, nil
	}
	return &ticketMessage, err
}

func (t *TicketMessages) CreateTicketMessage(ticketMessage *model.TicketMessage) (*model.TicketMessage, error) {
	_, err := t.DB.Model(ticketMessage).Returning("*").Insert()
	if len(ticketMessage.ID) < 1 {
		return nil, err
	}
	return ticketMessage, err
}
