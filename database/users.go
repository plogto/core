package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Users struct {
	Queries *db.Queries
}

func (u *Users) CreateUser(ctx context.Context, email, fullName string) (*db.User, error) {
	newUser := db.CreateUserParams{
		Email:          email,
		FullName:       fullName,
		Username:       util.RandomString(15),
		InvitationCode: util.RandomString(7),
	}

	user, err := u.Queries.CreateUser(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	user, err := u.Queries.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUserByInvitationCode(ctx context.Context, invitationCode string) (*db.User, error) {
	user, err := u.Queries.GetUserByInvitationCode(ctx, invitationCode)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	user, err := u.Queries.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUserByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := u.Queries.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUserByUsernameOrEmail(ctx context.Context, value string) (*db.User, error) {
	user, err := u.Queries.GetUserByUsernameOrEmail(ctx, value)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetUsersByUsernameOrFullNameAndPageInfo(ctx context.Context, value string, limit int32) (*model.Users, error) {
	var edges []*model.UsersEdge

	users, err := u.Queries.GetUsersByUsernameOrFullNameAndPageInfo(ctx, db.GetUsersByUsernameOrFullNameAndPageInfoParams{
		Lower: value,
		Limit: limit,
	})

	if err != nil {
		return nil, err
	}

	for _, value := range users {
		edges = append(edges, &model.UsersEdge{Node: &db.User{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	return &model.Users{
		Edges: edges,
	}, err
}

func (u *Users) UpdateUser(ctx context.Context, user *db.User) (*db.User, error) {
	user, err := u.Queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Bio:             user.Bio,
		FullName:        user.FullName,
		Background:      user.Background,
		Avatar:          user.Avatar,
		BackgroundColor: user.BackgroundColor,
		PrimaryColor:    user.PrimaryColor,
		IsPrivate:       user.IsPrivate,
	})

	if err != nil {
		return nil, err
	}

	return user, err
}
