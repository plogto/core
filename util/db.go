package util

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type Data interface {
	model.Tag |
		db.Password |
		db.Connection |
		db.Tag |
		db.File |
		db.Ticket |
		db.TicketMessage |
		db.CreditTransaction |
		db.CreditTransactionInfo |
		db.CreditTransactionTemplate |
		db.CreditTransactionDescriptionVariable |
		db.User |
		db.Post |
		db.PostTag |
		db.LikedPost |
		db.SavedPost |
		db.NotificationType |
		db.Notification
}

func HandleDBResponseWithoutError[T Data](data *T, err error) *T {
	if err != nil {
		return nil
	}

	return data
}

func HandleDBResponse[T Data](data *T, err error) (*T, error) {
	if err != nil {
		return nil, nil
	}

	return data, nil
}
