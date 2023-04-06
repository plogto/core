package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIsSavedPostExists(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsSavedPostExists(nil),
			Message:  "Should return false if savedPost is nil",
		},
		{
			Expected: false,
			Actual:   IsSavedPostExists(fixtures.EmptySavedPost),
			Message:  "Should return false if savedPost.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsSavedPostExists(fixtures.SavedPostWithID),
			Message:  "Should return true if savedPost is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
