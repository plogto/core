package validation

import (
	"testing"

	"github.com/plogto/core/graph/model"
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
			Actual:   IsPostExists(&model.Post{}),
			Message:  "Should return false if post.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsPostExists(&model.Post{ID: "id"}),
			Message:  "Should return true if post is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
	}
}
