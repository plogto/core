package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) CreateTicket(ctx context.Context, subject string, message string) (*model.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket := &model.Ticket{
		UserID:  user.ID,
		Subject: subject,
		Url:     util.RandomString(9),
	}

	s.Tickets.CreateTicket(ticket)

	if len(ticket.ID) > 0 {
		s.TicketMessages.CreateTicketMessage(&model.TicketMessage{
			SenderID: user.ID,
			TicketID: ticket.ID,
			Message:  message,
		})
	}

	return ticket, nil
}

func (s *Service) GetTicketByID(ctx context.Context, id string) (*model.Ticket, error) {
	return s.Tickets.GetTicketByID(id)
}

func (s *Service) GetTickets(ctx context.Context, input *model.PageInfoInput) (*model.Tickets, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	pageInfoInput := util.ExtractPageInfo(input)

	if user.Role == model.UserRoleAdmin {
		return s.Tickets.GetTicketsByUserIDAndPageInfo(&user.ID, *pageInfoInput.First, *pageInfoInput.After)
	}

	return s.Tickets.GetTicketsByUserIDAndPageInfo(&user.ID, *pageInfoInput.First, *pageInfoInput.After)
}
