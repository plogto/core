package model

import (
	"time"
)

type CreditTransaction struct {
	tableName               struct{} `pg:"credit_transactions"`
	ID                      string
	SenderID                string
	ReceiverID              string
	Amount                  float64
	Url                     string
	Description             *string
	Status                  CreditTransactionStatus
	CreditTransactionTypeID *string
	CreatedAt               *time.Time
	UpdatedAt               *time.Time
	DeletedAt               *time.Time `pg:"-,soft_delete"`
}
