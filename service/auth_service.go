package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
	"google.golang.org/api/idtoken"
)

func (s *Service) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := s.Users.GetUserByUsernameOrEmail(input.Username)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	password, err := s.Passwords.GetPasswordByUserID(user.ID)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	err = password.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	return s.PrepareAuthToken(user)
}

func (s *Service) Register(ctx context.Context, input model.RegisterInput, isOAuth bool) (*model.AuthResponse, error) {
	if _, err := s.Users.GetUserByUsernameOrEmail(input.Email); err == nil {
		return nil, errors.New("email has already been taken")
	}

	user := &model.User{
		Email:          input.Email,
		Username:       util.RandomString(15),
		InvitationCode: util.RandomString(7),
		FullName:       input.FullName,
	}

	newUser, err := s.Users.CreateUser(user)
	if err != nil {
		log.Printf("error while creating a user: %v", err)
		return nil, err
	}

	if !isOAuth {
		password := &model.Password{
			UserID: newUser.ID,
		}

		if err = password.HashPassword(input.Password); err != nil {
			log.Printf("error while hashing password: %v", err)
			return nil, errors.New("something went wrong")
		}

		if _, err := s.Passwords.AddPassword(password); err != nil {
			log.Printf("error white adding password: %v", err)
			return nil, err
		}
	}

	if input.InvitationCode != nil {
		if inviter, err := s.Users.GetUserByInvitationCode(*input.InvitationCode); err == nil {
			// FIXME
			InviterID, _ := uuid.Parse(inviter.ID)
			InviteeID, _ := uuid.Parse(newUser.ID)
			s.InvitedUsers.CreateInvitedUser(ctx, db.CreateInvitedUserParams{
				InviterID: InviterID,
				InviteeID: InviteeID,
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
			var ContentID uuid.UUID
			// FIXME
			ContentID, _ = uuid.Parse(newUser.ID)

			s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(ctx, db.CreateCreditTransactionDescriptionVariableParams{
				CreditTransactionInfoID: inviterTransactionCreditInfo.ID,
				Type:                    db.CreditTransactionDescriptionVariableTypeUser,
				Key:                     db.CreditTransactionDescriptionVariableKeyInvitedUser,
				ContentID:               ContentID,
			})

			inviteeTransactionCreditInfo, err := s.TransferCreditFromAdmin(ctx, TransferCreditFromAdminParams{
				ReceiverID:   newUser.ID,
				Status:       model.CreditTransactionStatusApproved,
				Type:         model.CreditTransactionTypeOrder,
				TemplateName: db.CreditTransactionTemplateNameRegisterByInvitationCode,
			})

			if err != nil {
				return nil, errors.New(err.Error())
			}

			ContentID, _ = uuid.Parse(inviter.ID)

			s.CreditTransactionDescriptionVariables.CreateCreditTransactionDescriptionVariable(ctx, db.CreateCreditTransactionDescriptionVariableParams{
				CreditTransactionInfoID: inviteeTransactionCreditInfo.ID,
				Type:                    db.CreditTransactionDescriptionVariableTypeUser,
				Key:                     db.CreditTransactionDescriptionVariableKeyInviterUser,
				ContentID:               ContentID,
			})
		}
	}

	plogAccount, _ := s.GetPlogAccount()
	if validation.IsUserExists(plogAccount) {
		// FIXME
		senderID, _ := uuid.Parse(plogAccount.ID)
		receiverID, _ := uuid.Parse(newUser.ID)
		s.CreateNotification(ctx, CreateNotificationArgs{
			Name:       db.NotificationTypeNameWelcome,
			SenderID:   senderID,
			ReceiverID: receiverID,
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

	user, _ := s.Users.GetUserByEmail(email)
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

func (s *Service) PrepareAuthToken(user *model.User) (*model.AuthResponse, error) {
	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
