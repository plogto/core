package validation

import (
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsPostExists(post *model.Post) bool {
	return post != nil && lo.IsNotEmpty(post.ID)
}
