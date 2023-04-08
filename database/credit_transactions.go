package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type CreditTransactions struct {
	Queries *db.Queries
}

func (c *CreditTransactions) CreateCreditTransaction(ctx context.Context, arg db.CreateCreditTransactionParams) (*db.CreditTransaction, error) {
	return util.HandleDBResponse(c.Queries.CreateCreditTransaction(ctx, arg))
}

func (c *CreditTransactions) GetCreditTransactionByID(ctx context.Context, id uuid.UUID) (*db.CreditTransaction, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionByID(ctx, id))
}

func (c *CreditTransactions) GetCreditTransactionByUrl(ctx context.Context, url string) (*db.CreditTransaction, error) {
	return util.HandleDBResponse(c.Queries.GetCreditTransactionByUrl(ctx, url))
}

func (c *CreditTransactions) GetCreditsByUserID(ctx context.Context, userID uuid.UUID) (float64, error) {
	amount, _ := c.Queries.GetCreditsByUserID(ctx, userID)

	return float64(amount), nil
}

func (c *CreditTransactions) GetCreditTransactionsByUserIDAndPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after time.Time) (*model.CreditTransactions, error) {
	var edges []*model.CreditTransactionsEdge
	var endCursor string

	creditTransactions, _ := c.Queries.GetCreditTransactionsByUserIDAndPageInfo(ctx, db.GetCreditTransactionsByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: after,
	})

	totalCount, _ := c.Queries.CountCreditTransactionsByUserIDAndPageInfo(ctx, db.CountCreditTransactionsByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
	})

	for _, value := range creditTransactions {
		edges = append(edges, &model.CreditTransactionsEdge{Node: &db.CreditTransaction{
			ID:        value.ID,
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

	return &model.CreditTransactions{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
	}, nil
}
