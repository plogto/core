package service

import (
	"context"
	"os"

	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

type CreateCreditTransactionInput struct {
	SenderID    string
	ReceiverID  string
	Amount      float64
	Description *string
	Status      model.CreditTransactionStatus
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

func (s *Service) CreateCreditTransaction(creditTransaction model.CreditTransaction) (*model.CreditTransaction, error) {
	receiver, _ := s.Users.GetUserByID(creditTransaction.ReceiverID)
	receiver.Credits = receiver.Credits + creditTransaction.Amount
	s.Users.UpdateUser(receiver)

	sender, _ := s.Users.GetUserByID(creditTransaction.SenderID)
	sender.Credits = sender.Credits - creditTransaction.Amount
	s.Users.UpdateUser(sender)

	creditTransaction.Url = util.RandomString(24)

	return s.CreditTransactions.CreateCreditTransaction(&creditTransaction)
}

func (s *Service) GetBankAccount() (*model.User, error) {
	username := os.Getenv("BANK_ACCOUNT")
	return s.Users.GetUserByUsername(username)
}

func (s *Service) GenerateCredits(amount float64) (*model.User, error) {
	bankUser, _ := s.GetBankAccount()

	if amount > 0 {
		bankUser.Credits += amount
		return s.Users.UpdateUser(bankUser)
	}

	return nil, nil
}

func (s *Service) TransferCreditFromAdmin(creditTransaction model.CreditTransaction) (*model.CreditTransaction, error) {
	s.GenerateCredits(creditTransaction.Amount)

	bankUser, _ := s.GetBankAccount()
	creditTransaction.SenderID = bankUser.ID

	return s.CreateCreditTransaction(creditTransaction)
}
