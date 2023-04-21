package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var TagID = pgtype.UUID{}
var EmptyTag = &db.Tag{}
var TagWithID = &db.Tag{ID: TagID}
var TagWithParentID = &db.Tag{ID: TagID}
