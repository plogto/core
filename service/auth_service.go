package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
	"google.golang.org/api/idtoken"
)

func (s *Service) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, _ := s.Users.GetUserByUsernameOrEmail(ctx, input.Username)
	if !validation.IsUserExists(user) {
		return nil, errors.New("username or password is not valid")
	}

	password, err := s.Passwords.GetPasswordByUserID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	err = util.ComparePassword(password.Password, input.Password)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	return s.PrepareAuthToken(user)
}

func (s *Service) Register(ctx context.Context, input model.RegisterInput, isOAuth bool) (*model.AuthResponse, error) {
	if user, _ := s.Users.GetUserByUsernameOrEmail(ctx, input.Email); validation.IsUser(user) {
		return nil, errors.New("email has already been taken")
	}

	user, err := s.Users.CreateUser(ctx, input.Email, input.FullName)
	if !validation.IsUser(user) {
		log.Printf("error while creating a user: %v", err)
		return nil, err
	}

	if !isOAuth {
		password := db.CreatePasswordParams{
			UserID: user.ID,
		}

		hashedPassword, err := util.HashPassword(input.Password)
		if err != nil {
			log.Printf("error while hashing password: %v", err)
			return nil, errors.New("something went wrong")
		}

		password.Password = *hashedPassword
		if _, err := s.Passwords.AddPassword(ctx, password); err != nil {
			log.Printf("error white adding password: %v", err)
			return nil, err
		}
	}

	if input.InvitationCode != nil {
		if inviter, err := s.Users.GetUserByInvitationCode(ctx, *input.InvitationCode); err == nil {
			s.InvitedUsers.CreateInvitedUser(ctx, db.CreateInvitedUserParams{
				InviterID: inviter.ID,
				InviteeID: user.ID,
			})

			if err != nil {
				return nil, errors.New(err.Error())
			}

			// transfer credits
			inviterTransactionCreditInfo, err := s.TransferCreditFromAdmin(ctx, TransferCreditFromAdminParams{
				ReceiverID:   inviter.ID,
				Status:       model.CreditTransactionStatusApproved,
				Type:         model.CreditTransactionTypeOrder,
				TemplateName: db.CreditTransactionTemplateNameInviteUser,
			})

			if err != nil {
				return nil, errors.New(err.Error())
			}

			s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(ctx, db.CreateCreditTransactionDescriptionVariableParams{
				CreditTransactionInfoID: inviterTransactionCreditInfo.ID,
				Type:                    db.CreditTransactionDescriptionVariableTypeUser,
				Key:                     db.CreditTransactionDescriptionVariableKeyInvitedUser,
				ContentID:               user.ID,
			})

			inviteeTransactionCreditInfo, err := s.TransferCreditFromAdmin(ctx, TransferCreditFromAdminParams{
				ReceiverID:   user.ID,
				Status:       model.CreditTransactionStatusApproved,
				Type:         model.CreditTransactionTypeOrder,
				TemplateName: db.CreditTransactionTemplateNameRegisterByInvitationCode,
			})

			if err != nil {
				return nil, errors.New(err.Error())
			}

			s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(ctx, db.CreateCreditTransactionDescriptionVariableParams{
				CreditTransactionInfoID: inviteeTransactionCreditInfo.ID,
				Type:                    db.CreditTransactionDescriptionVariableTypeUser,
				Key:                     db.CreditTransactionDescriptionVariableKeyInviterUser,
				ContentID:               inviter.ID,
			})
		}
	}

	plogAccount, _ := s.GetPlogAccount(ctx)
	if validation.IsUserExists(plogAccount) {
		s.CreateNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameWelcome,
			SenderID:   plogAccount.ID,
			ReceiverID: user.ID,
			Url:        "/" + plogAccount.Username,
		})
	}

	return s.PrepareAuthToken(user)
}

func (s *Service) OAuthGoogle(ctx context.Context, input model.OAuthGoogleInput) (*model.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), input.Credential, os.Getenv("GOOGLE_OAUTH_CLIENT_ID"))

	if err != nil {
		log.Printf("error while validating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	email := fmt.Sprintf("%v", payload.Claims["email"])

	user, _ := s.Users.GetUserByEmail(ctx, email)
	if s.PrepareUser(user) != nil {
		return s.PrepareAuthToken(user)
	}

	inputRegister := model.RegisterInput{
		FullName:       fmt.Sprintf("%v", payload.Claims["name"]),
		Email:          email,
		InvitationCode: input.InvitationCode,
		Password:       "",
	}

	return s.Register(ctx, inputRegister, true)
}

func (s *Service) PrepareAuthToken(user *db.User) (*model.AuthResponse, error) {
	token, err := util.GenToken(user.ID)
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
