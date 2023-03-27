package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var SavedPostID, _ = uuid.NewUUID()
var EmptySavedPost = &db.SavedPost{}
var SavedPostWithID = &db.SavedPost{ID: SavedPostID}
