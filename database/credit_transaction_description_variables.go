package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type CreditTransactionDescriptionVariables struct {
	DB *pg.DB
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByID(id string) (*model.CreditTransactionDescriptionVariable, error) {
	return c.GetCreditTransactionDescriptionVariableByField("id", id)
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByCreditTransactionID(creditTransactionID string) (*model.CreditTransactionDescriptionVariable, error) {
	return c.GetCreditTransactionDescriptionVariableByField("credit_transaction_id", creditTransactionID)
}

func (c *CreditTransactionDescriptionVariables) GetCreditTransactionDescriptionVariableByField(field string, value string) (*model.CreditTransactionDescriptionVariable, error) {
	var creditTransactionDescriptionVariable model.CreditTransactionDescriptionVariable
	err := c.DB.Model(&creditTransactionDescriptionVariable).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionDescriptionVariable, err
}

func (c *CreditTransactionDescriptionVariables) CreateCreditTransactionDescriptionVariable(creditTransactionDescriptionVariable *model.CreditTransactionDescriptionVariable) (*model.CreditTransactionDescriptionVariable, error) {
	_, err := c.DB.Model(creditTransactionDescriptionVariable).Returning("*").Insert()
	fmt.Println("CreateCreditTransactionDescriptionVariable", err)
	return creditTransactionDescriptionVariable, err
}
