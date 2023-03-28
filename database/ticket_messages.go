package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type TicketMessages struct {
	Queries *db.Queries
}

func (t *TicketMessages) CreateTicketMessage(ctx context.Context, arg db.CreateTicketMessageParams) (*db.TicketMessage, error) {
	ticketMessage, err := t.Queries.CreateTicketMessage(ctx, arg)

	if err != nil {
		return nil, err
	}

	return ticketMessage, nil
}

func (t *TicketMessages) GetTicketMessageByID(ctx context.Context, id uuid.UUID) (*db.TicketMessage, error) {
	ticketMessage, err := t.Queries.GetTicketMessageByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return ticketMessage, nil
}

func (t *TicketMessages) GetLastTicketMessageByTicketID(ctx context.Context, ticketID uuid.UUID) (*db.TicketMessage, error) {
	ticketMessage, err := t.Queries.GetLastTicketMessageByTicketID(ctx, ticketID)

	if err != nil {
		return nil, err
	}

	return ticketMessage, nil
}

func (t *TicketMessages) GetTicketMessagesByTicketIDAndPageInfo(ctx context.Context, ticketID uuid.UUID, limit int32, after string) (*model.TicketMessages, error) {
	var edges []*model.TicketMessagesEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	ticketMessages, err := t.Queries.GetTicketMessagesByTicketIDAndPageInfo(ctx, db.GetTicketMessagesByTicketIDAndPageInfoParams{
		TicketID:  ticketID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := t.Queries.CountTicketMessagesByTicketIDAndPageInfo(ctx, db.CountTicketMessagesByTicketIDAndPageInfoParams{
		TicketID:  ticketID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	for _, value := range ticketMessages {
		edges = append(edges, &model.TicketMessagesEdge{Node: &db.TicketMessage{
			ID:        value.ID,
			TicketID:  value.TicketID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.TicketMessages{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (t *TicketMessages) UpdateReadTicketMessagesByUserIDAndTicketID(ctx context.Context, userID string, ticketID uuid.UUID) (bool, error) {
	UserID, _ := uuid.Parse(userID)

	_, err := t.Queries.UpdateReadTicketMessagesByUserIDAndTicketID(ctx, db.UpdateReadTicketMessagesByUserIDAndTicketIDParams{
		SenderID: UserID,
		TicketID: ticketID,
	})

	if err != nil {
		return false, nil
	}

	return true, nil
}
