package validation

import "github.com/plogto/core/graph/model"

func IsPostExists(post *model.Post) bool {
	return len(post.ID) > 0
}
