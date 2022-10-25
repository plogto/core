package validation

import (
	"testing"

	"github.com/plogto/core/graph/model"
	"github.com/stretchr/testify/assert"
)

type SavedPostTestData struct {
	Expected, Actual bool
	Message          string
}

func TestIsSavedPostExists(t *testing.T) {
	var testData = []SavedPostTestData{
		{
			Expected: false,
			Actual:   IsSavedPostExists(nil),
			Message:  "Should return false if savedPost is nil",
		},
		{
			Expected: false,
			Actual:   IsSavedPostExists(&model.SavedPost{}),
			Message:  "Should return false if savedPost.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsSavedPostExists(&model.SavedPost{ID: "id"}),
			Message:  "Should return true if savedPost is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
	}
}
