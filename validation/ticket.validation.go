package validation

import "github.com/plogto/core/graph/model"

func IsTicketOpen(ticket *model.Ticket) bool {
	return len(ticket.ID) > 0 && ticket.Status == model.TicketStatusOpen
}

func IsTicketClosed(ticket *model.Ticket) bool {
	return len(ticket.ID) > 0 && ticket.Status == model.TicketStatusClosed
}

func IsTicketApproved(ticket *model.Ticket) bool {
	return len(ticket.ID) > 0 && ticket.Status == model.TicketStatusApproved
}

func IsUserAllowToUpdateTicket(user *model.User, ticket *model.Ticket) bool {
	return IsAdmin(user) || IsSuperAdmin(user) || user.ID == ticket.UserID
}

func IsTicketExist(ticket *model.Ticket) bool {
	return ticket == nil || len(ticket.ID) == 0
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
	case model.TicketStatusApproved:
		return model.TicketPermissionApprove
	case model.TicketStatusSolved:
		return model.TicketPermissionSolve
	default:
		return ""
	}
}
