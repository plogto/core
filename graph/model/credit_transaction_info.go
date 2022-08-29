package model

import (
	"time"
)

type CreditTransactionInfo struct {
	tableName                   struct{} `pg:"credit_transaction_infos"`
	ID                          string
	Description                 *string
	Status                      CreditTransactionStatus
	CreditTransactionTemplateID *string `pg:",use_zero"`
	CreatedAt                   *time.Time
	UpdatedAt                   *time.Time
	DeletedAt                   *time.Time `pg:"-,soft_delete"`
}
