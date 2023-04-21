package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var ConnectionID = pgtype.UUID{}
var ConnectionWithID = &db.Connection{ID: ConnectionID}
var ConnectionWithAcceptedStatus = &db.Connection{ID: ConnectionID, Status: 2}
var ConnectionWithPendingStatus = &db.Connection{ID: ConnectionID, Status: 1}
