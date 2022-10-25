package validation

import (
	"testing"

	"github.com/plogto/core/graph/model"
	"github.com/stretchr/testify/assert"
)

type LikedPostTestData struct {
	Expected, Actual bool
	Message          string
}

func TestIsLikedPostExists(t *testing.T) {
	var testData = []LikedPostTestData{
		{
			Expected: false,
			Actual:   IsLikedPostExists(nil),
			Message:  "Should return false if likedPost is nil",
		},
		{
			Expected: false,
			Actual:   IsLikedPostExists(&model.LikedPost{}),
			Message:  "Should return false if likedPost.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsLikedPostExists(&model.LikedPost{ID: "id"}),
			Message:  "Should return true if likedPost is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
	}
}
