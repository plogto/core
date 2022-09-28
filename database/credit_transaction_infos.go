package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type CreditTransactionInfos struct {
	DB *pg.DB
}

func (c *CreditTransactionInfos) CreateCreditTransactionInfo(creditTransactionInfo *model.CreditTransactionInfo) (*model.CreditTransactionInfo, error) {
	_, err := c.DB.Model(creditTransactionInfo).Returning("*").Insert()
	return creditTransactionInfo, err
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByID(id string) (*model.CreditTransactionInfo, error) {
	return c.GetCreditTransactionInfoByField("id", id)
}

func (c *CreditTransactionInfos) GetCreditTransactionInfoByField(field string, value string) (*model.CreditTransactionInfo, error) {
	var creditTransactionInfo model.CreditTransactionInfo
	err := c.DB.Model(&creditTransactionInfo).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionInfo, err
}

func (c *CreditTransactionInfos) UpdateCreditTransactionInfoStatus(creditTransactionInfo *model.CreditTransactionInfo) (*model.CreditTransactionInfo, error) {
	query := c.DB.Model(creditTransactionInfo).
		Where("id = ?id")

	_, err := query.Set("status = ?status").Returning("*").Update()

	return creditTransactionInfo, err
}
