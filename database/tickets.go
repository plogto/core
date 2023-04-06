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

func (t *Tickets) CreateTicket(ctx context.Context, userID uuid.UUID, subject string) (*db.Ticket, error) {
	newTicket := db.CreateTicketParams{
		Subject: subject,
		UserID:  userID,
		Url:     util.RandomHexString(9),
	}

	return util.HandleDBResponse(t.Queries.CreateTicket(ctx, newTicket))
}

func (t *Tickets) GetTicketByID(ctx context.Context, id uuid.UUID) (*db.Ticket, error) {
	return util.HandleDBResponse(t.Queries.GetTicketByID(ctx, id))
}

func (t *Tickets) GetTicketByURL(ctx context.Context, url string) (*db.Ticket, error) {
	return util.HandleDBResponse(t.Queries.GetTicketByURL(ctx, url))
}

func (t *Tickets) GetTicketsByUserIDAndPageInfo(ctx context.Context, userID uuid.NullUUID, limit int32, after time.Time) (*model.Tickets, error) {
	var edges []*model.TicketsEdge
	var endCursor string

	tickets, _ := t.Queries.GetTicketsByUserIDAndPageInfo(ctx, db.GetTicketsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		UpdatedAt: after,
	})

	totalCount, _ := t.Queries.CountTicketsByUserIDAndPageInfo(ctx, db.CountTicketsByUserIDAndPageInfoParams{
		UserID:    userID,
		UpdatedAt: after,
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
	}, nil
}

func (t *Tickets) UpdateTicketStatus(ctx context.Context, id uuid.UUID, status db.TicketStatusType) (*db.Ticket, error) {
	return util.HandleDBResponse(t.Queries.UpdateTicketStatus(ctx, db.UpdateTicketStatusParams{
		ID:     id,
		Status: status,
	}))
}

func (t *Tickets) UpdateTicketUpdatedAt(ctx context.Context, id uuid.UUID) (*db.Ticket, error) {
	UpdatedAt := time.Now()

	return util.HandleDBResponse(t.Queries.UpdateTicketUpdatedAt(ctx, db.UpdateTicketUpdatedAtParams{
		ID:        id,
		UpdatedAt: UpdatedAt,
	}))
}
