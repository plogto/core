package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var LikedPostID, _ = uuid.NewUUID()
var EmptyLikedPost = &db.LikedPost{}
var LikedPostWithID = &db.LikedPost{ID: LikedPostID}
