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
		Join("INNER JOIN credit_transaction_infos ON credit_transaction_infos.id = credit_transaction.credit_transaction_info_id").
		Where("credit_transaction.user_id = ?", userID).
		Where("credit_transaction.deleted_at is ?", nil).
		Order("credit_transaction_infos.created_at DESC")

	if len(after) > 0 {
		query.Where("credit_transaction_infos.created_at < ?", after)
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

func (c *CreditTransactions) GetCreditsByUserID(userID string) (float64, error) {
	var creditTransactions []*model.CreditTransaction

	err := c.DB.Model(&creditTransactions).
		ColumnExpr("sum(credit_transaction.amount) as amount").
		Join("INNER JOIN credit_transaction_infos ON credit_transaction_infos.id = credit_transaction.credit_transaction_info_id").
		Where("credit_transaction.user_id = ?", userID).
		Where("credit_transaction_infos.status = ?", model.CreditTransactionStatusApproved).
		Where("credit_transaction.deleted_at is ?", nil).
		Select()

	fmt.Println(err)

	return creditTransactions[0].Amount, err
}
