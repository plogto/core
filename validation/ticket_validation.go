package validation

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsTicketExist(ticket *db.Ticket) bool {
	return ticket != nil && lo.IsNotEmpty(ticket.ID)
}

func IsTicketOwner(user *db.User, ticket *db.Ticket) bool {
	return IsTicketExist(ticket) && IsUserExists(user) && user.ID == ticket.UserID
}

func IsUserAllowToUpdateTicket(user *db.User, ticket *db.Ticket) bool {
	return IsTicketOwner(user, ticket) || IsAdmin(user) || IsSuperAdmin(user)
}

func CheckUserPermission(permissions []*model.TicketPermission, status db.TicketStatusType) bool {
	isAllow := false

	for _, value := range permissions {
		if *value == ConvertTicketStatusToPermission(status) {
			isAllow = true
		}
	}

	return isAllow
}

func ConvertTicketStatusToPermission(status db.TicketStatusType) model.TicketPermission {
	switch status {
	case db.TicketStatusTypeOpen:
		return model.TicketPermissionOpen
	case db.TicketStatusTypeClosed:
		return model.TicketPermissionClose
	case db.TicketStatusTypeAccepted:
		return model.TicketPermissionAccept
	case db.TicketStatusTypeApproved:
		return model.TicketPermissionApprove
	case db.TicketStatusTypeRejected:
		return model.TicketPermissionReject
	case db.TicketStatusTypeSolved:
		return model.TicketPermissionSolve
	default:
		return ""
	}
}
