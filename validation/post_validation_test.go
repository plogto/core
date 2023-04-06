package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIsPostExists(t *testing.T) {
	var testData = []TestData{
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

func TestIsParentPostExists(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsParentPostExists(nil),
			Message:  "Should return false if post is nil",
		},
		{
			Expected: false,
			Actual:   IsParentPostExists(fixtures.EmptyPost),
			Message:  "Should return false if post.ParentID is not exist",
		},
		{
			Expected: true,
			Actual:   IsParentPostExists(fixtures.PostWithParentID),
			Message:  "Should return true if post.ParentID is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
