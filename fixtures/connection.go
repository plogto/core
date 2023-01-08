package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var ConnectionID, _ = uuid.NewUUID()
var ConnectionWithID = &db.Connection{ID: ConnectionID}
var ConnectionWithAcceptedStatus = &db.Connection{ID: ConnectionID, Status: 2}
var ConnectionWithPendingStatus = &db.Connection{ID: ConnectionID, Status: 1}
