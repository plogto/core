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
	Type                        db.CreditTransactionType
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

func (s *Service) GetCreditsByUserID(ctx context.Context, userID uuid.UUID) (float64, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	// FIXME
	if err != nil {
		return 0, nil
	}

	return s.CreditTransactions.GetCreditsByUserID(ctx, userID)
}

func (s *Service) GetCreditTransactionByID(ctx context.Context, id string) (*db.CreditTransaction, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	// FIXME
	ID, _ := uuid.Parse(id)
	return s.CreditTransactions.GetCreditTransactionByID(ctx, ID)
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

	senderID, _ := uuid.Parse(creditTransactionParams.SenderID)
	receiverID, _ := uuid.Parse(creditTransactionParams.ReceiverID)

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  senderID,
		RecipientID:             receiverID,
		Amount:                  float32(-amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  receiverID,
		RecipientID:             senderID,
		Amount:                  float32(amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	return creditTransactionInfo, err
}

func (s *Service) GetBankAccount() (*model.User, error) {
	username := os.Getenv("BANK_ACCOUNT")
	return s.Users.GetUserByUsername(username)
}

func (s *Service) GenerateCredits(ctx context.Context, amount sql.NullFloat64) (*db.CreditTransaction, error) {
	bankUser, _ := s.GetBankAccount()

	creditTransactionInfo, _ := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Status: db.CreditTransactionStatusApproved,
	})

	bankUserID, _ := uuid.Parse(bankUser.ID)

	return s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  bankUserID,
		RecipientID:             bankUserID,
		Amount:                  float32(amount.Float64),
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    db.CreditTransactionTypeFund,
		Url:                     util.RandomString(24),
	})
}

func (s *Service) TransferCreditFromAdmin(ctx context.Context, transferCreditFromAdminParams TransferCreditFromAdminParams) (*db.CreditTransactionInfo, error) {
	bankUser, _ := s.GetBankAccount()

	bankUserID, _ := uuid.Parse(bankUser.ID)

	bankUserCredits, _ := s.CreditTransactions.GetCreditsByUserID(ctx, bankUserID)

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
