package validation

import (
	"testing"

	"github.com/plogto/core/graph/model"
	"github.com/stretchr/testify/assert"
)

type UserTestData struct {
	Expected, Actual bool
	Message          string
}

var emptyUser = &model.User{}
var userWithID = &model.User{ID: "id"}
var userWithSuperAdminRole = &model.User{ID: "id", Role: model.UserRoleSuperAdmin}
var userWithAdminRole = &model.User{ID: "id", Role: model.UserRoleAdmin}
var userWithUserRole = &model.User{ID: "id", Role: model.UserRoleUser}

func TestIsUserExists(t *testing.T) {
	var testData = []UserTestData{
		{
			Expected: false,
			Actual:   IsUserExists(nil),
			Message:  "Should return false if user is nil",
		},
		{
			Expected: false,
			Actual:   IsUserExists(emptyUser),
			Message:  "Should return false if user.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsUserExists(userWithID),
			Message:  "Should return true if user is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
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
			Actual:   IsSuperAdmin(emptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsSuperAdmin(userWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsSuperAdmin(userWithSuperAdminRole),
			Message:  "Should return true if user.Role is SUPER_ADMIN",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
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
			Actual:   IsAdmin(emptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsAdmin(userWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsAdmin(userWithAdminRole),
			Message:  "Should return true if user.Role is ADMIN",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
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
			Actual:   IsUser(emptyUser),
			Message:  "Should return false if user.Role is not exist",
		},
		{
			Expected: false,
			Actual:   IsUser(userWithID),
			Message:  "Should return false if user.ID is exist but user.Role is not exist",
		},
		{
			Expected: true,
			Actual:   IsUser(userWithUserRole),
			Message:  "Should return true if user.Role is USER",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Actual, value.Expected, value.Message)
	}
}
