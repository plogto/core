package service

import (
	"context"

	"github.com/plogto/core/constants/err"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) CreateTicket(ctx context.Context, input model.CreateTicketInput) (*model.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket := &model.Ticket{
		UserID:  user.ID,
		Subject: input.Subject,
		Url:     util.RandomHexString(9),
	}

	s.Tickets.CreateTicket(ticket)

	if len(ticket.ID) > 0 {
		s.AddTicketMessage(ctx, ticket.ID, model.AddTicketMessageInput{
			Attachment: input.Attachment,
			Message:    input.Message,
		})
	}

	return ticket, nil
}

func (s *Service) GetTicketByID(ctx context.Context, id string) (*model.Ticket, error) {
	return s.Tickets.GetTicketByID(id)
}

func (s *Service) GetTickets(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.Tickets, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	pageInfo := util.ExtractPageInfo(pageInfoInput)

	if user.Role == model.UserRoleUser {
		return s.Tickets.GetTicketsByUserIDAndPageInfo(&user.ID, *pageInfo.First, *pageInfo.After)
	}

	return s.Tickets.GetTicketsByUserIDAndPageInfo(nil, *pageInfo.First, *pageInfo.After)
}

func (s *Service) CloseTicket(ticket model.Ticket) (*model.Ticket, error) {
	if validation.IsTicketOpen(&ticket) {
		ticket.Status = model.TicketStatusClosed
		closedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
		return closedTicket, nil
	}

	return nil, err.ErrorTicketIsNotOpen
}

func (s *Service) OpenTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	if validation.IsUser(&user) {
		return nil, err.ErrorAccessDenied
	}

	if validation.IsTicketClosed(&ticket) {
		ticket.Status = model.TicketStatusOpen
		openTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
		return openTicket, nil
	}

	return nil, err.ErrorTicketIsNotClosed
}

func (s *Service) ApproveTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	if validation.IsUser(&user) {
		return nil, err.ErrorAccessDenied
	}

	if validation.IsTicketOpen(&ticket) {
		ticket.Status = model.TicketStatusApproved
		approvedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
		return approvedTicket, nil
	}

	return nil, err.ErrorTicketIsNotOpen
}

func (s *Service) SolveTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	if validation.IsUser(&user) {
		return nil, err.ErrorAccessDenied
	}

	if validation.IsTicketApproved(&ticket) {
		ticket.Status = model.TicketStatusSolved
		solvedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
		return solvedTicket, nil
	}

	return nil, err.ErrorTicketIsNotApproved
}

func (s *Service) UpdateTicketStatus(ctx context.Context, ticketID string, status model.TicketStatus) (*model.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket, _ := s.Tickets.GetTicketByID(ticketID)

	if !validation.IsUserAllowToUpdateTicket(user, ticket) {
		return nil, err.ErrorAccessDenied
	}

	if validation.IsTicketExist(ticket) {
		return nil, err.ErrorTicketNotFound
	}

	switch status {
	case model.TicketStatusClosed:
		return s.CloseTicket(*ticket)
	case model.TicketStatusOpen:
		return s.OpenTicket(*user, *ticket)
	case model.TicketStatusApproved:
		return s.ApproveTicket(*user, *ticket)
	case model.TicketStatusSolved:
		return s.SolveTicket(*user, *ticket)
	}

	return nil, nil
}
