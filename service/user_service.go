package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	graph "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/middleware"
	"github.com/plogto/core/util"
)

func (s *Service) GetUserInfo(ctx context.Context) (*db.User, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	return s.PrepareUser(user), nil
}

func (s *Service) GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	user, _ := graph.GetUserLoader(ctx).Load(id.String())

	return s.PrepareUser(user), nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*db.User, error) {
	user, _ := s.Users.GetUserByUsername(ctx, username)

	return s.PrepareUser(user), nil
}

func (s *Service) GetUserByInvitationCode(ctx context.Context, invitationCode string) (*db.User, error) {
	user, _ := s.Users.GetUserByInvitationCode(ctx, invitationCode)

	return s.PrepareUser(user), nil
}

func (s *Service) SearchUser(ctx context.Context, expression string) (*model.Users, error) {
	var limit = constants.USERS_PAGE_LIMIT
	users, _ := s.Users.GetUsersByUsernameOrFullNameAndPageInfo(ctx, expression+"%", int32(limit))

	return users, nil
}

func (s *Service) CheckUserAccess(ctx context.Context, user *db.User, followingUser *db.User) bool {
	if followingUser.IsPrivate == bool(true) {
		if user != nil {
			connection, _ := s.Connections.GetConnection(ctx, followingUser.ID, user.ID)

			if followingUser.ID != user.ID {
				if len(connection.ID) < 1 || connection.Status < int32(2) {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

func (s *Service) EditUser(ctx context.Context, input model.EditUserInput) (*db.User, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	didUpdate := false

	if input.Username != nil {
		user.Username = *input.Username
		didUpdate = true
	}

	if input.BackgroundColor != nil {
		user.BackgroundColor = convertor.ModelBackgroundColorToDB(*input.BackgroundColor)
		didUpdate = true
	}

	if input.PrimaryColor != nil {
		user.PrimaryColor = convertor.ModelPrimaryColorToDB(*input.PrimaryColor)
		didUpdate = true
	}

	if input.Avatar != nil {
		if len(*input.Avatar) == 0 {
			user.Avatar = uuid.NullUUID{}
		} else {
			user.Avatar = uuid.NullUUID{uuid.MustParse(*input.Avatar), true}
		}
		didUpdate = true
	}

	if input.Background != nil {
		if len(*input.Background) == 0 {
			user.Background = uuid.NullUUID{}
		} else {
			user.Background = uuid.NullUUID{uuid.MustParse(*input.Background), true}
		}
		didUpdate = true
	}

	if input.FullName != nil {
		user.FullName = *input.FullName
		didUpdate = true
	}

	if input.Bio != nil {
		user.Bio = sql.NullString{*input.Bio, true}
		didUpdate = true
	}

	if input.Email != nil {
		user.Email = *input.Email
		didUpdate = true
	}

	if input.IsPrivate != nil {
		user.IsPrivate = *input.IsPrivate
		didUpdate = true
	}

	if !didUpdate {
		return nil, nil
	}

	updatedUser, _ := s.Users.UpdateUser(ctx, user)

	return s.PrepareUser(updatedUser), nil
}

func (s *Service) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.AuthResponse, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	password, _ := s.Passwords.GetPasswordByUserID(ctx, user.ID)

	if err = util.ComparePassword(password.Password, input.OldPassword); err != nil {
		return nil, errors.New("old password is not valid")
	}

	hashedPassword, err := util.HashPassword(input.NewPassword)

	if err != nil {
		password.Password = *hashedPassword
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	if _, err := s.Passwords.UpdatePassword(ctx, db.UpdatePasswordParams{
		UserID:   password.UserID,
		Password: *hashedPassword,
	}); err != nil {
		log.Printf("error white updating password: %v", err)
		return nil, err
	}

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

func (s *Service) CheckUsername(ctx context.Context, username string) (*db.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, _ := s.Users.GetUserByUsername(ctx, username)

	return s.PrepareUser(user), nil
}

func (s *Service) CheckEmail(ctx context.Context, email string) (*db.User, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, _ := s.Users.GetUserByEmail(ctx, email)

	return s.PrepareUser(user), nil
}

func (s *Service) PrepareUser(user *db.User) *db.User {
	if user == nil || len(user.ID) == 0 {
		return nil
	}

	return user
}

func (s *Service) GetPlogAccount(ctx context.Context) (*db.User, error) {
	username := os.Getenv("PLOG_ACCOUNT")
	return s.Users.GetUserByUsername(ctx, username)
}
