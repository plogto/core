package validation

import (
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsLikedPostExists(likedPost *model.LikedPost) bool {
	return likedPost != nil && lo.IsNotEmpty(likedPost.ID)
}
