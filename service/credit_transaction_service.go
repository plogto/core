package service

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateCreditTransactionParams struct {
	SenderID                    pgtype.UUID
	ReceiverID                  pgtype.UUID
	Amount                      pgtype.Float4
	Description                 pgtype.Text
	Status                      model.CreditTransactionStatus
	Type                        db.CreditTransactionType
	TemplateName                db.CreditTransactionTemplateName
	RelevantCreditTransactionID *string
}

type TransferCreditFromAdminParams struct {
	ReceiverID   pgtype.UUID
	Amount       *float64
	Status       model.CreditTransactionStatus
	Type         model.CreditTransactionType
	TemplateName db.CreditTransactionTemplateName
}

func (s *Service) GetCreditsByUserID(ctx context.Context, userID pgtype.UUID) (float64, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return 0, nil
	}

	return s.CreditTransactions.GetCreditsByUserID(ctx, userID)
}

func (s *Service) GetCreditTransactionByID(ctx context.Context, id pgtype.UUID) (*db.CreditTransaction, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	return s.CreditTransactions.GetCreditTransactionByID(ctx, id)
}

func (s *Service) GetCreditTransactions(ctx context.Context, pageInfo *model.PageInfoInput) (*model.CreditTransactions, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, nil
	}

	pagination := util.ExtractPageInfo(pageInfo)

	return s.CreditTransactions.GetCreditTransactionsByUserIDAndPageInfo(ctx, user.ID, pagination.First, pagination.After)
}

func (s *Service) CreateCreditTransaction(ctx context.Context, creditTransactionParams CreateCreditTransactionParams) (*db.CreditTransactionInfo, error) {
	var amount pgtype.Float4
	var creditTransactionTemplateID pgtype.UUID

	creditTransactionTemplate, _ := s.CreditTransactionTemplates.GetCreditTransactionTemplateByName(ctx, creditTransactionParams.TemplateName)
	creditTransactionTemplateID = creditTransactionTemplate.ID
	amount = creditTransactionTemplate.Amount

	if creditTransactionParams.Amount.Valid {
		amount = creditTransactionParams.Amount
	}

	creditTransactionInfo, err := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Description:                 creditTransactionParams.Description,
		CreditTransactionTemplateID: creditTransactionTemplateID,
		Status:                      convertor.ModelCreditTransactionStatusToDB(creditTransactionParams.Status),
	})

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  creditTransactionParams.SenderID,
		RecipientID:             creditTransactionParams.ReceiverID,
		Amount:                  -amount.Float32,
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  creditTransactionParams.ReceiverID,
		RecipientID:             creditTransactionParams.SenderID,
		Amount:                  amount.Float32,
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

func (s *Service) GenerateCredits(ctx context.Context, amount pgtype.Float4) (*db.CreditTransaction, error) {
	bankUser, _ := s.GetBankAccount(ctx)

	creditTransactionInfo, _ := s.CreditTransactionInfos.CreateCreditTransactionInfo(ctx, db.CreateCreditTransactionInfoParams{
		Status: db.CreditTransactionStatusApproved,
	})

	return s.CreditTransactions.CreateCreditTransaction(ctx, db.CreateCreditTransactionParams{
		UserID:                  bankUser.ID,
		RecipientID:             bankUser.ID,
		Amount:                  amount.Float32,
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

func (s *Service) GetDescriptionVariableContentByTypeAndContentID(ctx context.Context, creditTransactionDescriptionVariableType db.CreditTransactionDescriptionVariableType, contentID pgtype.UUID) (DescriptionVariable, error) {
	var descriptionVariable DescriptionVariable
	switch creditTransactionDescriptionVariableType {
	case db.CreditTransactionDescriptionVariableTypeUser:
		user, _ := graph.GetUserLoader(ctx).Load(convertor.UUIDToString(contentID))
		descriptionVariable = DescriptionVariable{
			Content: user.FullName,
			Url:     &user.Username,
			// Image: // implement GetFileLoader,
		}
	case db.CreditTransactionDescriptionVariableTypeTag:
		tag, _ := graph.GetTagLoader(ctx).Load(convertor.UUIDToString(contentID))
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
