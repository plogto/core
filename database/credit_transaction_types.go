package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type CreditTransactionTypes struct {
	DB *pg.DB
}

func (c *CreditTransactionTypes) GetCreditTransactionTypeByID(id string) (*model.CreditTransactionType, error) {
	return c.GetCreditTransactionTypeByField("id", id)
}

func (c *CreditTransactionTypes) GetCreditTransactionTypeByName(name model.CreditTransactionTypeName) (*model.CreditTransactionType, error) {
	var creditTransactionType model.CreditTransactionType
	err := c.DB.Model(&creditTransactionType).
		Where("name = ?", name).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionType, err
}

func (c *CreditTransactionTypes) GetCreditTransactionTypeByField(field string, value string) (*model.CreditTransactionType, error) {

	var creditTransactionType model.CreditTransactionType
	err := c.DB.Model(&creditTransactionType).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionType, err
}
