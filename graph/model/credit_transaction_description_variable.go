package model

import "time"

type CreditTransactionDescriptionVariable struct {
	tableName               struct{} `pg:"credit_transaction_description_variables"`
	ID                      string
	CreditTransactionInfoID string
	Type                    CreditTransactionDescriptionVariableType
	Key                     string
	ContentID               string
	DeletedAt               *time.Time `pg:"-,soft_delete"`
}
