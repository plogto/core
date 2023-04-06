package validation

import (
	"github.com/plogto/core/db"
	"github.com/samber/lo"
)

func IsLikedPostExists(likedPost *db.LikedPost) bool {
	return likedPost != nil && lo.IsNotEmpty(likedPost.ID)
}
