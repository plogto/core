package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsSavedPostExists(savedPost *db.SavedPost) bool {
	return savedPost != nil && lo.IsNotEmpty(savedPost.ID)
}
