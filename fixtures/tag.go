package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var TagID, _ = uuid.NewUUID()
var EmptyTag = &db.Tag{}
var TagWithID = &db.Tag{ID: TagID}
var TagWithParentID = &db.Tag{ID: TagID}
