package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsPostExists(post *db.Post) bool {
	return post != nil && lo.IsNotEmpty(post.ID)
}

func IsParentPostExists(post *db.Post) bool {
	return IsPostExists(post) && lo.IsNotEmpty(post.ParentID)
}
