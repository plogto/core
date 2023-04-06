package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) AddTicketMessage(ctx context.Context, ticketID uuid.UUID, input model.AddTicketMessageInput) (*db.TicketMessage, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	for _, id := range input.Attachment {
		file, _ := s.Files.GetFileByID(ctx, uuid.MustParse(*id))
		if file == nil {
			return nil, errors.New("attachment is not valid")
		}
	}

	ticketMessage, _ := s.TicketMessages.CreateTicketMessage(ctx, db.CreateTicketMessageParams{
		SenderID: user.ID,
		TicketID: ticketID,
		Message:  input.Message,
	})

	s.Tickets.UpdateTicketUpdatedAt(ctx, ticketID)

	if len(input.Attachment) > 0 {
		for _, v := range input.Attachment {
			s.TicketMessageAttachments.CreateTicketMessageAttachment(ctx, ticketMessage.ID, uuid.MustParse(*v))
		}
	}

	return ticketMessage, nil
}

func (s *Service) GetTicketMessageByID(ctx context.Context, id uuid.UUID) (*db.TicketMessage, error) {
	return s.TicketMessages.GetTicketMessageByID(ctx, id)
}

func (s *Service) GetLastTicketMessageByTicketID(ctx context.Context, ticketID uuid.UUID) (*db.TicketMessage, error) {
	return s.TicketMessages.GetLastTicketMessageByTicketID(ctx, ticketID)
}

func (s *Service) GetTicketMessagesByTicketURL(ctx context.Context, ticketURL string, pageInfo *model.PageInfoInput) (*model.TicketMessages, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if !validation.IsUserExists(user) {
		return nil, nil
	}

	pagination := util.ExtractPageInfo(pageInfo)

	ticket, _ := s.Tickets.GetTicketByURL(ctx, ticketURL)

	if validation.IsUser(user) && user.ID != ticket.UserID {
		return nil, nil
	}

	return s.TicketMessages.GetTicketMessagesByTicketIDAndPageInfo(ctx, ticket.ID, int32(pagination.First), pagination.After)
}

func (s *Service) ReadTicketMessages(ctx context.Context, ticketID uuid.UUID) (*bool, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	if user == nil {
		return nil, nil
	}

	ticket, _ := s.Tickets.GetTicketByID(ctx, ticketID)

	status, _ := s.TicketMessages.UpdateReadTicketMessagesByUserIDAndTicketID(ctx, user.ID, ticket.ID)

	return &status, nil
}
