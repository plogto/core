package validation

import (
	"testing"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/fixtures"
	"github.com/plogto/core/graph/model"
	"github.com/stretchr/testify/assert"
)

type TicketTestData struct {
	Expected, Actual bool
	Message          string
}

type TicketPermissionTestData struct {
	Expected, Actual model.TicketPermission
	Message          string
}

func TestIsTicketExists(t *testing.T) {
	var testData = []TicketTestData{
		{
			Expected: false,
			Actual:   IsTicketExist(nil),
			Message:  "Should return false if ticket is nil",
		},
		{
			Expected: false,
			Actual:   IsTicketExist(fixtures.EmptyTicket),
			Message:  "Should return false if ticket.ID is not exist",
		},
		{
			Expected: true,
			Actual:   IsTicketExist(fixtures.TicketWithID),
			Message:  "Should return true if ticket is exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsTicketOwner(t *testing.T) {
	var testData = []TicketTestData{
		{
			Expected: false,
			Actual:   IsTicketOwner(nil, nil),
			Message:  "Should return false if ticket and user are nil",
		},
		{
			Expected: false,
			Actual:   IsTicketOwner(fixtures.EmptyUser, fixtures.EmptyTicket),
			Message:  "Should return false if ticket.ID and user.ID are not exist",
		},
		{
			Expected: true,
			Actual:   IsTicketOwner(fixtures.UserWithID, fixtures.TicketWithUserID),
			Message:  "Should return true if ticket and user are exist",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestIsUserAllowToUpdateTicket(t *testing.T) {
	var testData = []TicketTestData{
		{
			Expected: false,
			Actual:   IsUserAllowToUpdateTicket(nil, nil),
			Message:  "Should return false if ticket and user are nil",
		},
		{
			Expected: false,
			Actual:   IsUserAllowToUpdateTicket(fixtures.EmptyUser, fixtures.EmptyTicket),
			Message:  "Should return false if ticket.ID and user.ID are not exist",
		},
		{
			Expected: true,
			Actual:   IsUserAllowToUpdateTicket(fixtures.UserWithID, fixtures.TicketWithUserID),
			Message:  "Should return true if ticket and user are exist",
		},
		{
			Expected: true,
			Actual:   IsUserAllowToUpdateTicket(fixtures.UserWithAdminRole, fixtures.TicketWithUserID),
			Message:  "Should return true if ticket is exist and user.Role is ADMIN",
		},
		{
			Expected: true,
			Actual:   IsUserAllowToUpdateTicket(fixtures.UserWithAdminRole, fixtures.TicketWithUserID),
			Message:  "Should return true if ticket is exist and user.Role is SUPER_ADMIN",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestCheckUserPermission(t *testing.T) {
	var permissions []*model.TicketPermission

	var testData = []TicketTestData{
		{
			Expected: false,
			Actual:   CheckUserPermission(nil, model.TicketStatusOpen),
			Message:  "Should return false if permission is nil",
		},
		{
			Expected: false,
			Actual: CheckUserPermission(append(
				permissions,
				&constants.CLOSE,
			), model.TicketStatusOpen),
			Message: "Should return false if ticket status is not included in the permissions",
		},
		{
			Expected: true,
			Actual: CheckUserPermission(append(
				permissions,
				&constants.OPEN,
			), model.TicketStatusOpen),
			Message: "Should return false if ticket status is included in the permissions",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}

func TestConvertTicketStatusToPermission(t *testing.T) {
	var testData = []TicketPermissionTestData{
		{
			Expected: "",
			Actual:   ConvertTicketStatusToPermission(model.TicketStatus("")),
			Message:  "Should return empty string if ticket status is empty",
		},
		{
			Expected: model.TicketPermissionOpen,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusOpen),
			Message:  "Should return OPEN if ticket status is OPEN",
		},
		{
			Expected: model.TicketPermissionClose,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusClosed),
			Message:  "Should return CLOSE if ticket status is OPENED",
		},
		{
			Expected: model.TicketPermissionAccept,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusAccepted),
			Message:  "Should return ACCEPT if ticket status is ACCEPTED",
		},
		{
			Expected: model.TicketPermissionApprove,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusApproved),
			Message:  "Should return APPROVE if ticket status is APPROVED",
		},
		{
			Expected: model.TicketPermissionReject,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusRejected),
			Message:  "Should return REJECT if ticket status is REJECTED",
		},
		{
			Expected: model.TicketPermissionSolve,
			Actual:   ConvertTicketStatusToPermission(model.TicketStatusSolved),
			Message:  "Should return SOLVE if ticket status is SOLVED",
		},
	}

	for _, value := range testData {
		assert.Equal(t, value.Expected, value.Actual, value.Message)
	}
}
