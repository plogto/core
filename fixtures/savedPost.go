package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var SavedPostID = pgtype.UUID{}
var EmptySavedPost = &db.SavedPost{}
var SavedPostWithID = &db.SavedPost{ID: SavedPostID}
