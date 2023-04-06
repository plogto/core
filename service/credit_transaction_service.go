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
	SenderID                    uuid.UUID
	ReceiverID                  uuid.UUID
	Amount                      sql.NullFloat64
	Description                 *string
	Status                      model.CreditTransactionStatus
	Type                        db.CreditTransactionType
	TemplateName                db.CreditTransactionTemplateName
	RelevantCreditTransactionID *string
}

type TransferCreditFromAdminParams struct {
	ReceiverID   uuid.UUID
	Amount       *float64
	Status       model.CreditTransactionStatus
	Type         model.CreditTransactionType
	TemplateName db.CreditTransactionTemplateName
}

func (s *Service) GetCreditsByUserID(ctx context.Context, userID uuid.UUID) (float64, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return 0, nil
	}

	return s.CreditTransactions.GetCreditsByUserID(ctx, userID)
}

func (s *Service) GetCreditTransactionByID(ctx context.Context, id uuid.UUID) (*db.CreditTransaction, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactions.GetCreditTransactionByID(ctx, id)
}

func (s *Service) GetCreditTransactions(ctx context.Context, input *model.PageInfoInput) (*model.CreditTransactions, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	pageInfoInput := util.ExtractPageInfo(input)

	return s.CreditTransactions.GetCreditTransactionsByUserIDAndPageInfo(ctx, user.ID, int32(pageInfoInput.First), pageInfoInput.After)
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

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  creditTransactionParams.SenderID,
		RecipientID:             creditTransactionParams.ReceiverID,
		Amount:                  float32(-amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  creditTransactionParams.ReceiverID,
		RecipientID:             creditTransactionParams.SenderID,
		Amount:                  float32(amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	return creditTransactionInfo, err
}

func (s *Service) GetBankAccount(ctx context.Context) (*db.User, error) {
	username := os.Getenv("BANK_ACCOUNT")
	return s.Users.GetUserByUsername(ctx, username)
}

func (s *Service) GenerateCredits(ctx context.Context, amount sql.NullFloat64) (*db.CreditTransaction, error) {
	bankUser, _ := s.GetBankAccount(ctx)

	creditTransactionInfo, _ := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Status: db.CreditTransactionStatusApproved,
	})

	return s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  bankUser.ID,
		RecipientID:             bankUser.ID,
		Amount:                  float32(amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    db.CreditTransactionTypeFund,
		Url:                     util.RandomString(24),
	})
}

func (s *Service) TransferCreditFromAdmin(ctx context.Context, transferCreditFromAdminParams TransferCreditFromAdminParams) (*db.CreditTransactionInfo, error) {
	bankUser, _ := s.GetBankAccount(ctx)

	bankUserCredits, _ := s.CreditTransactions.GetCreditsByUserID(ctx, bankUser.ID)

	if bankUserCredits <= 0 {
		s.GenerateCredits(ctx, constants.GENERATE_CREDITS_AMOUNT)
	}

	return s.CreateCreditTransaction(ctx, CreateCreditTransactionParams{
		SenderID:     bankUser.ID,
		ReceiverID:   transferCreditFromAdminParams.ReceiverID,
		Status:       transferCreditFromAdminParams.Status,
		TemplateName: transferCreditFromAdminParams.TemplateName,
		Type:         db.CreditTransactionTypeOrder,
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
			// Image: // implement GetFileLoader,
		}
	case db.CreditTransactionDescriptionVariableTypeTag:
		tag, _ := graph.GetTagLoader(ctx).Load(contentID.String())
		descriptionVariable = DescriptionVariable{
			Content: tag.Name,
			Url:     &tag.Name,
		}
	case db.CreditTransactionDescriptionVariableTypeTicket:
		ticket, _ := s.Tickets.GetTicketByID(ctx, contentID)
		descriptionVariable = DescriptionVariable{
			Content: ticket.Subject,
			Url:     &ticket.Url,
		}
	}
	return descriptionVariable, nil
}
