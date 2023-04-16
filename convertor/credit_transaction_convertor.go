package convertor

import (
	"strings"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func ModelCreditTransactionStatusToDB(status model.CreditTransactionStatus) db.CreditTransactionStatus {
	return db.CreditTransactionStatus((strings.ToLower(status.String())))
}
