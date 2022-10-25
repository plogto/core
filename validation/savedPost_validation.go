package validation

import (
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsSavedPostExists(savedPost *model.SavedPost) bool {
	return savedPost != nil && lo.IsNotEmpty(savedPost.ID)
}
