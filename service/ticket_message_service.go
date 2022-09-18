package service

import (
	"context"
	"errors"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) AddTicketMessage(ctx context.Context, ticketID string, input model.AddTicketMessageInput) (*model.TicketMessage, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	for _, id := range input.Attachment {
		file, _ := s.Files.GetFileByID(*id)
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	ticketMessage, _ := s.TicketMessages.CreateTicketMessage(&model.TicketMessage{
		SenderID: user.ID,
		TicketID: ticketID,
		Message:  input.Message,
	})

	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			s.TicketMessageAttachments.CreateTicketMessageAttachment(&model.TicketMessageAttachment{
				TicketMessageID: ticketMessage.ID,
				FileID:          *v,
			})
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

	return s.TicketMessages.GetTicketMessagesByTicketIDAndPageInfo(ticket.ID, *pageInfo.First, *pageInfo.After)
}
