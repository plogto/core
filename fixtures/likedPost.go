package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var LikedPostID = pgtype.UUID{}
var EmptyLikedPost = &db.LikedPost{}
var LikedPostWithID = &db.LikedPost{ID: LikedPostID}
