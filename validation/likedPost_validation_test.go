package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIsLikedPostExists(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsLikedPostExists(nil),
			Message:  "Should return false if likedPost is nil",
		},
		{
			Expected: false,
			Actual:   IsLikedPostExists(fixtures.EmptyLikedPost),
			Message:  "Should return false if likedPost.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsLikedPostExists(fixtures.LikedPostWithID),
			Message:  "Should return true if likedPost is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
