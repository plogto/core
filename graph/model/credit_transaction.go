package model

import (
	"database/sql"
	"time"
)

type CreditTransaction struct {
	tableName                   struct{} `pg:"credit_transactions"`
	ID                          string
	CreditTransactionInfoID     string
	RelevantCreditTransactionID *string `pg:",use_zero"`
	UserID                      string
	RecipientID                 string
	Amount                      sql.NullFloat64
	Url                         string
	Type                        CreditTransactionType
	CreatedAt                   *time.Time
	DeletedAt                   *time.Time `pg:"-,soft_delete"`
}
