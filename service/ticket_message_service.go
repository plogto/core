package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddTicketMessage(ctx context.Context, ticketID string, input model.AddTicketMessageInput) (*model.TicketMessage, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	for _, id := range input.Attachment {
		ID, _ := uuid.Parse(*id)
		file, _ := s.Files.GetFileByID(ctx, ID)
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	ticketMessage, _ := s.TicketMessages.CreateTicketMessage(&model.TicketMessage{
		SenderID: user.ID,
		TicketID: ticketID,
		Message:  input.Message,
	})

	s.Tickets.UpdateTicketUpdatedAt(ticketID)

	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			V, _ := uuid.Parse(*v)
			TicketMessageID, _ := uuid.Parse(ticketMessage.ID)
			s.TicketMessageAttachments.CreateTicketMessageAttachment(ctx, TicketMessageID, V)
		}
	}

	return ticketMessage, nil
}

func (s *Service) GetTicketMessageByID(ctx context.Context, id string) (*model.TicketMessage, error) {
	return s.TicketMessages.GetTicketMessageByID(id)
}

func (s *Service) GetLastTicketMessageByTicketID(ctx context.Context, ticketID string) (*model.TicketMessage, error) {
	return s.TicketMessages.GetLastTicketMessageByTicketID(ticketID)
}

func (s *Service) GetTicketMessagesByTicketURL(ctx context.Context, ticketURL string, pageInfoInput *model.PageInfoInput) (*model.TicketMessages, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	pageInfo := util.ExtractPageInfo(pageInfoInput)

	ticket, _ := s.Tickets.GetTicketByURL(ticketURL)

	if user.Role == model.UserRoleUser && user.ID != ticket.UserID {
		return nil, nil
	}

	return s.TicketMessages.GetTicketMessagesByTicketIDAndPageInfo(ticket.ID, pageInfo.First, pageInfo.After)
}

func (s *Service) ReadTicketMessages(ctx context.Context, ticketID string) (*bool, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	ticket, _ := s.Tickets.GetTicketByID(ticketID)

	status, _ := s.TicketMessages.UpdateReadTicketMessagesByUserIDAndTicketID(user.ID, ticket.ID)

	return &status, nil
}
