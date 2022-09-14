package service

import (
	"context"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddTicketMessage(ctx context.Context, ticketID string, message string) (*model.TicketMessage, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	ticketMessage, _ := s.TicketMessages.CreateTicketMessage(&model.TicketMessage{
		SenderID: user.ID,
		TicketID: ticketID,
		Message:  message,
	})

	return ticketMessage, nil
}

func (s *Service) GetTicketMessageByID(ctx context.Context, id string) (*model.TicketMessage, error) {
	return s.TicketMessages.GetTicketMessageByID(id)
}

func (s *Service) GetTicketMessagesByTicketURL(ctx context.Context, ticketURL string, input *model.PageInfoInput) (*model.TicketMessages, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)

	ticket, _ := s.Tickets.GetTicketByURL(ticketURL)

	return s.TicketMessages.GetTicketMessagesByTicketIDAndPageInfo(ticket.ID, *pageInfoInput.First, *pageInfoInput.After)
}
