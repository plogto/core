package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type CreditTransactionTemplates struct {
	DB *pg.DB
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByID(id string) (*model.CreditTransactionTemplate, error) {
	return c.GetCreditTransactionTemplateByField("id", id)
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByName(name model.CreditTransactionTemplateName) (*model.CreditTransactionTemplate, error) {
	var creditTransactionTemplate model.CreditTransactionTemplate
	err := c.DB.Model(&creditTransactionTemplate).
		Where("name = ?", name).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionTemplate, err
}

func (c *CreditTransactionTemplates) GetCreditTransactionTemplateByField(field string, value string) (*model.CreditTransactionTemplate, error) {

	var creditTransactionTemplate model.CreditTransactionTemplate
	err := c.DB.Model(&creditTransactionTemplate).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &creditTransactionTemplate, err
}
