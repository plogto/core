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
