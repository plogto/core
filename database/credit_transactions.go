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
	creditTransaction, err := c.Queries.CreateCreditTransaction(ctx, arg)

	if err != nil {
		return nil, err
	}

	return creditTransaction, err
}

func (c *CreditTransactions) GetCreditTransactionByID(ctx context.Context, id uuid.UUID) (*db.CreditTransaction, error) {
	creditTransaction, err := c.Queries.GetCreditTransactionByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return creditTransaction, err
}

func (c *CreditTransactions) GetCreditTransactionByUrl(ctx context.Context, url string) (*db.CreditTransaction, error) {
	creditTransaction, err := c.Queries.GetCreditTransactionByUrl(ctx, url)

	if err != nil {
		return nil, err
	}

	return creditTransaction, err
}

func (c *CreditTransactions) GetCreditsByUserID(ctx context.Context, userID uuid.UUID) (float64, error) {
	amount, err := c.Queries.GetCreditsByUserID(ctx, userID)

	if err != nil {
		return 0, err
	}

	return float64(amount), err
}

func (c *CreditTransactions) GetCreditTransactionsByUserIDAndPageInfo(ctx context.Context, userID string, limit int32, after string) (*model.CreditTransactions, error) {
	var edges []*model.CreditTransactionsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)
	// FIXME
	UserID, _ := uuid.Parse(userID)

	creditTransactions, err := c.Queries.GetCreditTransactionsByUserIDAndPageInfo(ctx, db.GetCreditTransactionsByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    UserID,
		CreatedAt: createdAt,
	})

	totalCount, _ := c.Queries.CountCreditTransactionsByUserIDAndPageInfo(ctx, db.CountCreditTransactionsByUserIDAndPageInfoParams{
		UserID:    UserID,
		CreatedAt: createdAt,
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
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}
