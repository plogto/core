package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
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
		s.TicketMessages.CreateTicketMessage(&model.TicketMessage{
			SenderID: user.ID,
			TicketID: ticket.ID,
			Message:  input.Message,
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

	if user.Role == model.UserRoleAdmin {
		return s.Tickets.GetTicketsByUserIDAndPageInfo(nil, *pageInfo.First, *pageInfo.After)
	}

	return s.Tickets.GetTicketsByUserIDAndPageInfo(&user.ID, *pageInfo.First, *pageInfo.After)
}
