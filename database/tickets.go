package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Tickets struct {
	DB *pg.DB
}

func (t *Tickets) GetTicketByField(field string, value string) (*model.Ticket, error) {
	var ticket model.Ticket
	err := t.DB.Model(&ticket).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(ticket.ID) < 1 {
		return nil, nil
	}
	return &ticket, err
}

func (t *Tickets) GetTicketsByUserIDAndPageInfo(userID *string, limit int, after string) (*model.Tickets, error) {
	var tickets []*model.Ticket
	var edges []*model.TicketsEdge
	var endCursor string

	query := t.DB.Model(&tickets).
		Where("user_id = ?", userID).
		Where("deleted_at is ?", nil).
		Order("updated_at DESC")

	if len(after) > 0 {
		query.Where("updated_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range tickets {
		edges = append(edges, &model.TicketsEdge{Node: &model.Ticket{
			ID:        value.ID,
			UpdatedAt: value.UpdatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.UpdatedAt)
	}

	return &model.Tickets{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (t *Tickets) GetTicketByID(id string) (*model.Ticket, error) {
	return t.GetTicketByField("id", id)
}

func (t *Tickets) GetTicketByURL(url string) (*model.Ticket, error) {
	return t.GetTicketByField("url", url)
}

func (t *Tickets) CreateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	_, err := t.DB.Model(ticket).Returning("*").Insert()
	if len(ticket.ID) < 1 {
		return nil, err
	}
	return ticket, err
}
