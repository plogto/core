package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

type PostTestData struct {
	Expected, Actual bool
	Message          string
}

func TestIsPostExists(t *testing.T) {
	var testData = []PostTestData{
		{
			Expected: false,
			Actual:   IsPostExists(nil),
			Message:  "Should return false if post is nil",
		},
		{
			Expected: false,
			Actual:   IsPostExists(fixtures.EmptyPost),
			Message:  "Should return false if post.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsPostExists(fixtures.PostWithID),
			Message:  "Should return true if post is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
