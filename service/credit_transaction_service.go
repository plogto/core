package service

import (
	"context"
	"database/sql"
	"os"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateCreditTransactionParams struct {
	SenderID                    string
	ReceiverID                  string
	Amount                      sql.NullFloat64
	Description                 *string
	Status                      model.CreditTransactionStatus
	Type                        model.CreditTransactionType
	TemplateName                db.CreditTransactionTemplateName
	RelevantCreditTransactionID *string
}

type TransferCreditFromAdminParams struct {
	ReceiverID   string
	Amount       *float64
	Status       model.CreditTransactionStatus
	Type         model.CreditTransactionType
	TemplateName db.CreditTransactionTemplateName
}

func (s *Service) GetCreditsByUserID(ctx context.Context, userID *string) (float64, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if userID == nil || err != nil {
		return 0, nil
	}

	return s.CreditTransactions.GetCreditsByUserID(*userID)
}

func (s *Service) GetCreditTransactionByID(ctx context.Context, id *string) (*model.CreditTransaction, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if id == nil || err != nil {
		return nil, nil
	}

	return s.CreditTransactions.GetCreditTransactionByID(*id)
}

func (s *Service) GetCreditTransactions(ctx context.Context, input *model.PageInfoInput) (*model.CreditTransactions, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)

	return s.CreditTransactions.GetCreditTransactionsByUserIDAndPageInfo(user.ID, pageInfoInput.First, pageInfoInput.After)
}

func (s *Service) CreateCreditTransaction(ctx context.Context, creditTransactionParams CreateCreditTransactionParams) (*db.CreditTransactionInfo, error) {
	var amount sql.NullFloat64
	var creditTransactionTemplateID uuid.UUID

	creditTransactionTemplate, _ := s.CreditTransactionTemplates.GetCreditTransactionTemplateByName(ctx, creditTransactionParams.TemplateName)
	creditTransactionTemplateID = creditTransactionTemplate.ID
	amount = creditTransactionTemplate.Amount

	if creditTransactionParams.Amount.Valid {
		amount = creditTransactionParams.Amount
	}

	creditTransactionInfo, err := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Description:                 sql.NullString{*creditTransactionParams.Description, true},
		CreditTransactionTemplateID: uuid.NullUUID{creditTransactionTemplateID, true},
		Status:                      db.CreditTransactionStatus(creditTransactionParams.Status),
	})

	s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  creditTransactionParams.SenderID,
		RecipientID:             creditTransactionParams.ReceiverID,
		Amount:                  sql.NullFloat64{-amount.Float64, true},
		CreditTransactionInfoID: creditTransactionInfo.ID.String(),
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  creditTransactionParams.ReceiverID,
		RecipientID:             creditTransactionParams.SenderID,
		Amount:                  amount,
		CreditTransactionInfoID: creditTransactionInfo.ID.String(),
		Type:                    model.CreditTransactionType(creditTransactionParams.Type),
		Url:                     util.RandomString(24),
	})

	return creditTransactionInfo, err
}

func (s *Service) GetBankAccount() (*model.User, error) {
	username := os.Getenv("BANK_ACCOUNT")
	return s.Users.GetUserByUsername(username)
}

func (s *Service) GenerateCredits(ctx context.Context, amount sql.NullFloat64) (*model.CreditTransaction, error) {
	bankUser, _ := s.GetBankAccount()

	creditTransactionInfo, _ := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Status: db.CreditTransactionStatusApproved,
	})

	return s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  bankUser.ID,
		RecipientID:             bankUser.ID,
		Amount:                  amount,
		CreditTransactionInfoID: creditTransactionInfo.ID.String(),
		Type:                    model.CreditTransactionTypeFund,
		Url:                     util.RandomString(24),
	})
}

func (s *Service) TransferCreditFromAdmin(ctx context.Context, transferCreditFromAdminParams TransferCreditFromAdminParams) (*db.CreditTransactionInfo, error) {
	bankUser, _ := s.GetBankAccount()

	bankUserCredits, _ := s.CreditTransactions.GetCreditsByUserID(bankUser.ID)

	if bankUserCredits <= 0 {
		s.GenerateCredits(ctx, constants.GENERATE_CREDITS_AMOUNT)
	}

	return s.CreateCreditTransaction(ctx, CreateCreditTransactionParams{
		SenderID:     bankUser.ID,
		ReceiverID:   transferCreditFromAdminParams.ReceiverID,
		Status:       transferCreditFromAdminParams.Status,
		TemplateName: transferCreditFromAdminParams.TemplateName,
		Type:         model.CreditTransactionTypeOrder,
	})
}

func (s *Service) GetDescriptionVariableContentByTypeAndContentID(ctx context.Context, creditTransactionDescriptionVariableType db.CreditTransactionDescriptionVariableType, contentID uuid.UUID) (DescriptionVariable, error) {
	var descriptionVariable DescriptionVariable
	switch creditTransactionDescriptionVariableType {
	case db.CreditTransactionDescriptionVariableTypeUser:
		user, _ := graph.GetUserLoader(ctx).Load(contentID.String())
		descriptionVariable = DescriptionVariable{
			Content: user.FullName,
			Url:     &user.Username,
			Image:   user.Avatar,
		}
	case db.CreditTransactionDescriptionVariableTypeTag:
		tag, _ := graph.GetTagLoader(ctx).Load(contentID.String())
		descriptionVariable = DescriptionVariable{
			Content: tag.Name,
			Url:     &tag.Name,
		}
	case db.CreditTransactionDescriptionVariableTypeTicket:
		ticket, _ := s.Tickets.GetTicketByID(contentID.String())
		descriptionVariable = DescriptionVariable{
			Content: ticket.Subject,
			Url:     &ticket.Url,
		}
	}
	return descriptionVariable, nil
}
