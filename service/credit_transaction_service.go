package service

import (
	"context"
	"os"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateCreditTransactionParams struct {
	SenderID                    string
	ReceiverID                  string
	Amount                      *float64
	Description                 *string
	Status                      model.CreditTransactionStatus
	Type                        model.CreditTransactionType
	TemplateName                *model.CreditTransactionTemplateName
	RelevantCreditTransactionID *string
}

type TransferCreditFromAdminParams struct {
	ReceiverID   string
	Amount       *float64
	Status       model.CreditTransactionStatus
	Type         model.CreditTransactionType
	TemplateName model.CreditTransactionTemplateName
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

	return s.CreditTransactions.GetCreditTransactionsByUserIDAndPageInfo(user.ID, *pageInfoInput.First, *pageInfoInput.After)
}

func (s *Service) CreateCreditTransaction(creditTransactionParams CreateCreditTransactionParams) (*model.CreditTransactionInfo, error) {
	var amount float64
	var creditTransactionTemplateID string

	if creditTransactionParams.TemplateName != nil {
		creditTransactionTemplate, _ := s.CreditTransactionTemplates.GetCreditTransactionTemplateByName(*creditTransactionParams.TemplateName)
		creditTransactionTemplateID = creditTransactionTemplate.ID
		amount = *creditTransactionTemplate.Amount
	}

	if creditTransactionParams.Amount != nil {
		amount = *creditTransactionParams.Amount
	}

	creditTransactionInfo, err := s.CreditTransactionInfos.CreateCreditTransactionInfo(&model.CreditTransactionInfo{
		Description:                 creditTransactionParams.Description,
		CreditTransactionTemplateID: &creditTransactionTemplateID,
		Status:                      creditTransactionParams.Status,
	})

	s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  creditTransactionParams.SenderID,
		RecipientID:             creditTransactionParams.ReceiverID,
		Amount:                  -amount,
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    creditTransactionParams.Type,
		Url:                     util.RandomString(24),
	})

	s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  creditTransactionParams.ReceiverID,
		RecipientID:             creditTransactionParams.SenderID,
		Amount:                  amount,
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

func (s *Service) GenerateCredits(amount float64) (*model.CreditTransaction, error) {
	bankUser, _ := s.GetBankAccount()

	creditTransactionInfo, _ := s.CreditTransactionInfos.CreateCreditTransactionInfo(&model.CreditTransactionInfo{
		Status: model.CreditTransactionStatusApproved,
	})

	return s.CreditTransactions.CreateCreditTransaction(&model.CreditTransaction{
		UserID:                  bankUser.ID,
		RecipientID:             bankUser.ID,
		Amount:                  amount,
		CreditTransactionInfoID: creditTransactionInfo.ID,
		Type:                    model.CreditTransactionTypeFund,
		Url:                     util.RandomString(24),
	})
}

func (s *Service) TransferCreditFromAdmin(transferCreditFromAdminParams TransferCreditFromAdminParams) (*model.CreditTransactionInfo, error) {
	bankUser, _ := s.GetBankAccount()

	bankUserCredits, _ := s.CreditTransactions.GetCreditsByUserID(bankUser.ID)

	if bankUserCredits <= 0 {
		s.GenerateCredits(100)
	}

	return s.CreateCreditTransaction(CreateCreditTransactionParams{
		SenderID:     bankUser.ID,
		ReceiverID:   transferCreditFromAdminParams.ReceiverID,
		Status:       model.CreditTransactionStatusApproved,
		TemplateName: &transferCreditFromAdminParams.TemplateName,
		Type:         model.CreditTransactionTypeOrder,
	})
}

func (s *Service) GetDescriptionVariableContentByTypeAndContentID(ctx context.Context, creditTransactionDescriptionVariableType model.CreditTransactionDescriptionVariableType, contentID string) (DescriptionVariable, error) {
	var descriptionVariable DescriptionVariable
	switch creditTransactionDescriptionVariableType {
	case model.CreditTransactionDescriptionVariableTypeUser:
		user, _ := s.Users.GetUserByID(contentID)
		descriptionVariable = DescriptionVariable{
			Content: user.FullName,
			Url:     &user.Username,
			Image:   user.Avatar,
		}
	case model.CreditTransactionDescriptionVariableTypeTag:
		tag, _ := s.Tags.GetTagByID(contentID)
		descriptionVariable = DescriptionVariable{
			Content: tag.Name,
			Url:     &tag.Name,
		}
	}
	return descriptionVariable, nil
}
