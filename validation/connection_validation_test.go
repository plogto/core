package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIsConnectionExists(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsConnectionExists(nil),
			Message:  "Should return false if connection is nil",
		},
		{
			Expected: true,
			Actual:   IsConnectionExists(fixtures.ConnectionWithID),
			Message:  "Should return true if connection is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsConnectionStatusAccepted(t *testing.T) {
	var testData = []TestData{
		{
			Expected: false,
			Actual:   IsConnectionStatusAccepted(nil),
			Message:  "Should return false if connection is nil",
		},
		{
			Expected: false,
			Actual:   IsConnectionStatusAccepted(fixtures.ConnectionWithPendingStatus),
			Message:  "Should return false if connection status is pending",
		},
		{
			Expected: true,
			Actual:   IsConnectionStatusAccepted(fixtures.ConnectionWithAcceptedStatus),
			Message:  "Should return true if connection status is accepted",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
