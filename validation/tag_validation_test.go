package validation

import (
	"testing"

	"github.com/plogto/core/db"
	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIsTagExists(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsTagExists[db.Tag](nil),
			Message:  "Should return false if tag is nil",
		},
		{
			Expected: false,
			Actual:   IsTagExists(fixtures.EmptyTag),
			Message:  "Should return false if tag.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsTagExists(fixtures.TagWithID),
			Message:  "Should return true if tag is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
