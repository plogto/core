package service

import (
	"context"
	"errors"

	"github.com/plogto/core/constants"
	"github.com/plogto/core/constants/e"
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

func (s *Service) GetTicketPermissionsByTicketID(ctx context.Context, ticketID string) ([]*model.TicketPermission, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket, _ := s.Tickets.GetTicketByID(ticketID)
	var permissions []*model.TicketPermission

	if !validation.IsUserAllowToUpdateTicket(user, ticket) {
		return nil, e.ErrorAccessDenied
	}

	switch ticket.Status {
	case model.TicketStatusOpen:
		return s.GetPermissionsForOpenTicket(*user, *ticket)
	case model.TicketStatusClosed:
		return s.GetPermissionsForClosedTicket(*user)
	case model.TicketStatusApproved:
		return s.GetPermissionsForApprovedTicket(*user)
	case model.TicketStatusSolved:
		return s.GetPermissionsForSolvedTicket(*user)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForOpenTicket(user model.User, ticket model.Ticket) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(permissions, &constants.NEW_MESSAGE, &constants.APPROVE, &constants.CLOSE)
	case model.UserRoleAdmin:
		if ticket.UserID != user.ID {
			permissions = append(permissions, &constants.NEW_MESSAGE, &constants.APPROVE, &constants.CLOSE)
		} else {
			permissions = append(permissions, &constants.NEW_MESSAGE, &constants.CLOSE)
		}
	case model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE, &constants.CLOSE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForClosedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin, model.UserRoleAdmin:
		permissions = append(permissions, &constants.NEW_MESSAGE, &constants.OPEN)
	case model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForApprovedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(permissions, &constants.NEW_MESSAGE, &constants.SOLVE, &constants.CLOSE)
	case model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForSolvedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	case model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) CloseTicket(ticket model.Ticket) (*model.Ticket, error) {
	ticket.Status = model.TicketStatusClosed
	closedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
	return closedTicket, nil
}

func (s *Service) OpenTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	ticket.Status = model.TicketStatusOpen
	openTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
	return openTicket, nil
}

func (s *Service) ApproveTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	ticket.Status = model.TicketStatusApproved
	approvedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)

	transactionCreditInfo, _ := s.TransferCreditFromAdmin(TransferCreditFromAdminParams{
		ReceiverID:   ticket.UserID,
		Status:       model.CreditTransactionStatusPending,
		Type:         model.CreditTransactionTypeOrder,
		TemplateName: model.CreditTransactionTemplateNameApproveTicket,
	})

	s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(&model.CreditTransactionDescriptionVariable{
		CreditTransactionInfoID: transactionCreditInfo.ID,
		Type:                    model.CreditTransactionDescriptionVariableTypeTicket,
		Key:                     "ticket",
		ContentID:               ticket.ID,
	})

	return approvedTicket, nil
}

func (s *Service) SolveTicket(user model.User, ticket model.Ticket) (*model.Ticket, error) {
	ticket.Status = model.TicketStatusSolved
	solvedTicket, _ := s.Tickets.UpdateTicketStatus(&ticket)
	return solvedTicket, nil
}

func (s *Service) UpdateTicketStatus(ctx context.Context, ticketID string, status model.TicketStatus) (*model.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket, _ := s.Tickets.GetTicketByID(ticketID)

	if validation.IsTicketExist(ticket) {
		return nil, e.ErrorTicketNotFound
	}

	permissions, err := s.GetTicketPermissionsByTicketID(ctx, ticketID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if !validation.CheckUserPermission(permissions, status) {
		return nil, e.ErrorAccessDenied
	}

	switch status {
	case model.TicketStatusOpen:
		return s.OpenTicket(*user, *ticket)
	case model.TicketStatusClosed:
		return s.CloseTicket(*ticket)
	case model.TicketStatusApproved:
		return s.ApproveTicket(*user, *ticket)
	case model.TicketStatusSolved:
		return s.SolveTicket(*user, *ticket)
	}

	return nil, nil
}
