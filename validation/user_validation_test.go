package validation

import (
	"testing"

	"github.com/plogto/core/fixtures"
	"github.com/stretchr/testify/assert"
)

type UserTestData struct {
	Expected, Actual bool
	Message          string
}

func TestIsUserExists(t *testing.T) {
	var testData = []UserTestData{
		{
			Expected: false,
			Actual:   IsUserExists(nil),
			Message:  "Should return false if user is nil",
		},
		{
			Expected: false,
			Actual:   IsUserExists(fixtures.EmptyUser),
			Message:  "Should return false if user.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsUserExists(fixtures.UserWithID),
			Message:  "Should return true if user is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsSuperAdmin(t *testing.T) {
	var testData = []UserTestData{
		{
			Expected: false,
			Actual:   IsSuperAdmin(nil),
			Message:  "Should return false if user is nil",
		},
		{
			Expected: false,
			Actual:   IsSuperAdmin(fixtures.EmptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsSuperAdmin(fixtures.UserWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsSuperAdmin(fixtures.UserWithSuperAdminRole),
			Message:  "Should return true if user.Role is SUPER_ADMIN",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsAdmin(t *testing.T) {
	var testData = []UserTestData{
		{
			Expected: false,
			Actual:   IsAdmin(nil),
			Message:  "Should return false if user is nil",
		},
		{
			Expected: false,
			Actual:   IsAdmin(fixtures.EmptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsAdmin(fixtures.UserWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsAdmin(fixtures.UserWithAdminRole),
			Message:  "Should return true if user.Role is ADMIN",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsUser(t *testing.T) {
	var testData = []UserTestData{
		{
			Expected: false,
			Actual:   IsUser(nil),
			Message:  "Should return false if user is nil",
		},
		{
			Expected: false,
			Actual:   IsUser(fixtures.EmptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsUser(fixtures.UserWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsUser(fixtures.UserWithUserRole),
			Message:  "Should return true if user.Role is USER",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
