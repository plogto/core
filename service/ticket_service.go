package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/constants/e"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

func (s *Service) CreateTicket(ctx context.Context, input model.CreateTicketInput) (*db.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	ticket, _ := s.Tickets.CreateTicket(ctx, user.ID, input.Subject)

	if validation.IsTicketExist(ticket) {
		s.AddTicketMessage(ctx, ticket.ID, model.AddTicketMessageInput{
			Attachment: input.Attachment,
			Message:    input.Message,
		})
	}

	return ticket, nil
}

func (s *Service) GetTicketByID(ctx context.Context, id uuid.UUID) (*db.Ticket, error) {
	return s.Tickets.GetTicketByID(ctx, id)
}

func (s *Service) GetTickets(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.Tickets, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	pagination := util.ExtractPageInfo(pageInfoInput)

	if validation.IsUser(user) {
		return s.Tickets.GetTicketsByUserIDAndPageInfo(ctx, &user.ID, int32(pagination.First), pagination.After)
	}

	return s.Tickets.GetTicketsByUserIDAndPageInfo(ctx, nil, int32(pagination.First), pagination.After)
}

func (s *Service) GetTicketPermissionsByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*model.TicketPermission, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket, _ := s.Tickets.GetTicketByID(ctx, ticketID)
	var permissions []*model.TicketPermission

	if !validation.IsUserAllowToUpdateTicket(user, ticket) {
		return nil, e.ErrorAccessDenied
	}

	switch ticket.Status {
	case db.TicketStatusTypeOpen:
		return s.GetPermissionsForOpenTicket(*user, *ticket)
	case db.TicketStatusTypeClosed:
		return s.GetPermissionsForClosedTicket(*user)
	case db.TicketStatusTypeAccepted:
		return s.GetPermissionsForAcceptedTicket(*user)
	case db.TicketStatusTypeApproved:
		return s.GetPermissionsForApprovedTicket(*user)
	case db.TicketStatusTypeRejected:
		return s.GetPermissionsForRejectedTicket(*user)
	case db.TicketStatusTypeSolved:
		return s.GetPermissionsForSolvedTicket(*user)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForOpenTicket(user model.User, ticket db.Ticket) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(
			permissions,
			&constants.NEW_MESSAGE,
			&constants.ACCEPT,
			&constants.APPROVE,
			&constants.CLOSE,
		)
	case model.UserRoleAdmin:
		if ticket.UserID.String() != user.ID {
			permissions = append(permissions, &constants.NEW_MESSAGE, &constants.ACCEPT, &constants.CLOSE)
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

func (s *Service) GetPermissionsForAcceptedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(
			permissions,
			&constants.NEW_MESSAGE,
			&constants.APPROVE,
			&constants.REJECT,
			&constants.CLOSE,
		)
	case model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForApprovedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin:
		permissions = append(permissions, &constants.NEW_MESSAGE, &constants.SOLVE)
	case model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForRejectedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin, model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) GetPermissionsForSolvedTicket(user model.User) ([]*model.TicketPermission, error) {
	var permissions []*model.TicketPermission

	switch user.Role {
	case model.UserRoleSuperAdmin, model.UserRoleAdmin, model.UserRoleUser:
		permissions = append(permissions, &constants.NEW_MESSAGE)
	}

	return permissions, nil
}

func (s *Service) CloseTicket(ctx context.Context, ticket db.Ticket) (*db.Ticket, error) {
	closedTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeClosed)

	return closedTicket, nil
}

func (s *Service) OpenTicket(ctx context.Context, user model.User, ticket db.Ticket) (*db.Ticket, error) {
	openTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeOpen)
	return openTicket, nil
}

func (s *Service) AcceptTicket(ctx context.Context, user model.User, ticket db.Ticket) (*db.Ticket, error) {
	acceptedTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeAccepted)

	transactionCreditInfo, _ := s.TransferCreditFromAdmin(ctx, TransferCreditFromAdminParams{
		ReceiverID:   ticket.UserID.String(),
		Status:       model.CreditTransactionStatusPending,
		Type:         model.CreditTransactionTypeOrder,
		TemplateName: db.CreditTransactionTemplateNameApproveTicket,
	})

	s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(ctx, db.CreateCreditTransactionDescriptionVariableParams{
		CreditTransactionInfoID: transactionCreditInfo.ID,
		Type:                    db.CreditTransactionDescriptionVariableTypeTicket,
		Key:                     "ticket",
		ContentID:               ticket.ID,
	})

	return acceptedTicket, nil
}

func (s *Service) ApproveTicket(ctx context.Context, user model.User, ticket db.Ticket) (*db.Ticket, error) {
	oldStatus := ticket.Status

	if oldStatus != db.TicketStatusTypeAccepted {
		s.AcceptTicket(ctx, user, ticket)
	}

	approvedTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeApproved)
	// FIXME
	creditTransactionDescriptionVariables, _ := s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariableByContentID(ctx, ticket.ID)
	creditTransactionInfo := db.CreditTransactionInfo{
		// FIXME
		ID:     creditTransactionDescriptionVariables.CreditTransactionInfoID,
		Status: db.CreditTransactionStatusApproved,
	}

	s.CreditTransactionInfos.UpdateCreditTransactionInfoStatus(ctx, db.UpdateCreditTransactionInfoStatusParams{
		ID:     creditTransactionInfo.ID,
		Status: creditTransactionInfo.Status,
	})

	return approvedTicket, nil
}

func (s *Service) RejectTicket(ctx context.Context, user model.User, ticket db.Ticket) (*db.Ticket, error) {
	oldStatus := ticket.Status
	rejectedTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeRejected)

	if oldStatus == db.TicketStatusTypeAccepted {
		creditTransactionDescriptionVariables, _ := s.CreditTransactionDescriptionVariables.GetCreditTransactionDescriptionVariableByContentID(ctx, ticket.ID)
		creditTransactionInfo := db.CreditTransactionInfo{
			// FIXME
			ID:     creditTransactionDescriptionVariables.CreditTransactionInfoID,
			Status: db.CreditTransactionStatusCanceled,
		}

		s.CreditTransactionInfos.UpdateCreditTransactionInfoStatus(ctx, db.UpdateCreditTransactionInfoStatusParams{
			ID:     creditTransactionInfo.ID,
			Status: creditTransactionInfo.Status,
		})
	}

	return rejectedTicket, nil
}

func (s *Service) SolveTicket(ctx context.Context, user model.User, ticket db.Ticket) (*db.Ticket, error) {
	solvedTicket, _ := s.Tickets.UpdateTicketStatus(ctx, ticket.ID, db.TicketStatusTypeSolved)
	return solvedTicket, nil
}

func (s *Service) UpdateTicketStatus(ctx context.Context, ticketID uuid.UUID, status db.TicketStatusType) (*db.Ticket, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	ticket, _ := s.Tickets.GetTicketByID(ctx, ticketID)

	if !validation.IsTicketExist(ticket) {
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
	case db.TicketStatusTypeOpen:
		return s.OpenTicket(ctx, *user, *ticket)
	case db.TicketStatusTypeClosed:
		return s.CloseTicket(ctx, *ticket)
	case db.TicketStatusTypeAccepted:
		return s.AcceptTicket(ctx, *user, *ticket)
	case db.TicketStatusTypeApproved:
		return s.ApproveTicket(ctx, *user, *ticket)
	case db.TicketStatusTypeRejected:
		return s.RejectTicket(ctx, *user, *ticket)
	case db.TicketStatusTypeSolved:
		return s.SolveTicket(ctx, *user, *ticket)
	}

	return nil, nil
}
