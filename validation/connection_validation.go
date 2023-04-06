package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsConnectionExists(connection *db.Connection) bool {
	return connection != nil && lo.IsNotEmpty(connection.ID)
}

func IsConnectionStatusAccepted(connection *db.Connection) bool {
	return IsConnectionExists(connection) && connection.Status == 2
}
