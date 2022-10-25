package validation

import (
	"github.com/plogto/core/graph/model"
	"github.com/samber/lo"
)

func IsTicketExist(ticket *model.Ticket) bool {
	return ticket != nil && lo.IsNotEmpty(ticket.ID)
}

func IsTicketOpen(ticket *model.Ticket) bool {
	return IsTicketExist(ticket) && ticket.Status == model.TicketStatusOpen
}

func IsTicketClosed(ticket *model.Ticket) bool {
	return IsTicketExist(ticket) && ticket.Status == model.TicketStatusClosed
}

func IsTicketAccepted(ticket *model.Ticket) bool {
	return IsTicketExist(ticket) && ticket.Status == model.TicketStatusAccepted
}

func IsTicketApproved(ticket *model.Ticket) bool {
	return IsTicketExist(ticket) && ticket.Status == model.TicketStatusApproved
}

func IsTicketRejected(ticket *model.Ticket) bool {
	return IsTicketExist(ticket) && ticket.Status == model.TicketStatusRejected
}

func IsUserAllowToUpdateTicket(user *model.User, ticket *model.Ticket) bool {
	return IsAdmin(user) || IsSuperAdmin(user) || user.ID == ticket.UserID
}

func CheckUserPermission(permissions []*model.TicketPermission, status model.TicketStatus) bool {
	isAllow := false

	for _, value := range permissions {
		if *value == ConvertTicketStatusToPermission(status) {
			isAllow = true
		}
	}

	return isAllow
}

func ConvertTicketStatusToPermission(status model.TicketStatus) model.TicketPermission {
	switch status {
	case model.TicketStatusOpen:
		return model.TicketPermissionOpen
	case model.TicketStatusClosed:
		return model.TicketPermissionClose
	case model.TicketStatusAccepted:
		return model.TicketPermissionAccept
	case model.TicketStatusApproved:
		return model.TicketPermissionApprove
	case model.TicketStatusRejected:
		return model.TicketPermissionReject
	case model.TicketStatusSolved:
		return model.TicketPermissionSolve
	default:
		return ""
	}
}
