package validation

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type PostType interface {
	db.Post | model.Post
}

// TODO: fix validation
func IsPostExists[T PostType](post *T) bool {
	return post != nil
}

// TODO: fix validation
func IsParentPostExists(post *model.Post) bool {
	return IsPostExists(post) && post.ParentID.Valid
}
