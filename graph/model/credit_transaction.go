package model

import "time"

type CreditTransaction struct {
	tableName                   struct{} `pg:"credit_transactions"`
	ID                          string
	CreditTransactionInfoID     string
	RelevantCreditTransactionID *string `pg:",use_zero"`
	UserID                      string
	RecipientID                 string
	Amount                      float64
	Url                         string
	Type                        CreditTransactionType
	CreatedAt                   *time.Time
	DeletedAt                   *time.Time `pg:"-,soft_delete"`
}
