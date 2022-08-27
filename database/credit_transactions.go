package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type CreditTransactions struct {
	DB *pg.DB
}

func (c *CreditTransactions) CreateCreditTransaction(creditTransaction *model.CreditTransaction) (*model.CreditTransaction, error) {
	_, err := c.DB.Model(creditTransaction).Returning("*").Insert()
	fmt.Println("CreateCreditTransaction", err)
	return creditTransaction, err
}

func (c *CreditTransactions) GetCreditTransactionByID(id string) (*model.CreditTransaction, error) {
	return c.GetCreditTransactionByField("id", id)
}

func (c *CreditTransactions) GetCreditTransactionByUrl(url string) (*model.CreditTransaction, error) {
	return c.GetCreditTransactionByField("url", url)
}

func (c *CreditTransactions) GetCreditTransactionByField(field string, value string) (*model.CreditTransaction, error) {
	var creditTransaction model.CreditTransaction
	err := c.DB.Model(&creditTransaction).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransaction, err
}

func (c *CreditTransactions) GetCreditTransactionsByUserIDAndPageInfo(userID string, limit int, after string) (*model.CreditTransactions, error) {
	var creditTransactions []*model.CreditTransaction
	var edges []*model.CreditTransactionsEdge
	var endCursor string

	query := c.DB.Model(&creditTransactions).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.Where("sender_id = ?", userID).
				WhereOr("receiver_id = ?", false)
			return q, nil
		}).
		Where("deleted_at is ?", nil).
		Order("created_at DESC")

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range creditTransactions {
		edges = append(edges, &model.CreditTransactionsEdge{Node: &model.CreditTransaction{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > limit {
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
