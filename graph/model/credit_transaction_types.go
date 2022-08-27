package model

import (
	"time"
)

type CreditTransactionType struct {
	tableName struct{} `pg:"credit_transaction_types"`
	ID        string
	Name      CreditTransactionTypeName
	Template  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `pg:"-,soft_delete"`
}
