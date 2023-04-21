package fixtures

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

var PostID = pgtype.UUID{}
var EmptyPost = &db.Post{}
var PostWithID = &db.Post{ID: PostID}
var PostWithParentID = &db.Post{ID: PostID, ParentID: PostID}
