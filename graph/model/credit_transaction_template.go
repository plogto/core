package model

import "time"

type CreditTransactionTemplate struct {
	tableName struct{} `pg:"credit_transaction_templates"`
	ID        string
	Name      CreditTransactionTemplateName
	Template  string
	Amount    *float64   `pg:",use_zero"`
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
