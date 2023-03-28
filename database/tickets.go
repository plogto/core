package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Tickets struct {
	Queries *db.Queries
}

func (t *Tickets) CreateTicket(ctx context.Context, userID, subject string) (*db.Ticket, error) {
	UserID, _ := uuid.Parse(userID)

	newTicket := db.CreateTicketParams{
		Subject: subject,
		UserID:  UserID,
		Url:     util.RandomHexString(9),
	}

	ticket, err := t.Queries.CreateTicket(ctx, newTicket)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *Tickets) GetTicketByID(ctx context.Context, id uuid.UUID) (*db.Ticket, error) {
	ticket, err := t.Queries.GetTicketByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *Tickets) GetTicketByURL(ctx context.Context, url string) (*db.Ticket, error) {
	ticket, err := t.Queries.GetTicketByURL(ctx, url)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *Tickets) GetTicketsByUserIDAndPageInfo(ctx context.Context, userID *string, limit int32, after string) (*model.Tickets, error) {
	var edges []*model.TicketsEdge
	var endCursor string
	var UserID uuid.NullUUID

	updatedAt, _ := time.Parse(time.RFC3339, after)
	// FIXME
	if userID != nil {
		id, _ := uuid.Parse(*userID)
		UserID = uuid.NullUUID{id, true}
	}

	tickets, err := t.Queries.GetTicketsByUserIDAndPageInfo(ctx, db.GetTicketsByUserIDAndPageInfoParams{
		UserID:    UserID,
		Limit:     limit,
		UpdatedAt: updatedAt,
	})

	totalCount, _ := t.Queries.CountTicketsByUserIDAndPageInfo(ctx, db.CountTicketsByUserIDAndPageInfoParams{
		UserID:    UserID,
		Limit:     limit,
		UpdatedAt: updatedAt,
	})

	for _, value := range tickets {
		edges = append(edges, &model.TicketsEdge{Node: &db.Ticket{
			ID:        value.ID,
			UpdatedAt: value.UpdatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.UpdatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Tickets{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (t *Tickets) UpdateTicketStatus(ctx context.Context, id uuid.UUID, status db.TicketStatusType) (*db.Ticket, error) {
	ticket, err := t.Queries.UpdateTicketStatus(ctx, db.UpdateTicketStatusParams{
		ID:     id,
		Status: status,
	})

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (t *Tickets) UpdateTicketUpdatedAt(ctx context.Context, id uuid.UUID) (*db.Ticket, error) {
	UpdatedAt := time.Now()

	ticket, err := t.Queries.UpdateTicketUpdatedAt(ctx, db.UpdateTicketUpdatedAtParams{
		ID:        id,
		UpdatedAt: UpdatedAt,
	})

	return ticket, err
}
