package model

import (
	"time"
)

type CreditTransactionDescriptionVariable struct {
	tableName           struct{} `pg:"credit_transaction_description_variables"`
	ID                  string
	CreditTransactionID string
	Type                CreditTransactionDescriptionVariableType
	Key                 string
	ContentID           string
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	DeletedAt           *time.Time `pg:"-,soft_delete"`
}
