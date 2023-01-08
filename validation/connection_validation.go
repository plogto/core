package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsConnectionExists(connection *db.Connection) bool {
	return lo.IsNotEmpty(connection)
}

func IsConnectionStatusAccepted(connection *db.Connection) bool {
	return IsConnectionExists(connection) && connection.Status == 2
}
