package fixtures

import (
	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

var PostID, _ = uuid.NewUUID()
var EmptyPost = &db.Post{}
var PostWithID = &db.Post{ID: PostID}
var PostWithParentID = &db.Post{ID: PostID, ParentID: uuid.NullUUID{PostID, true}}
