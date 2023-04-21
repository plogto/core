package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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

	return u.Queries.CreateUser(ctx, newUser)
}

func (u *Users) GetUserByID(ctx context.Context, id pgtype.UUID) (*db.User, error) {
	user, err := u.Queries.GetUserByID(ctx, id)
	return user, err
}

func (u *Users) GetUserByInvitationCode(ctx context.Context, invitationCode string) (*db.User, error) {
	return u.Queries.GetUserByInvitationCode(ctx, invitationCode)
}

func (u *Users) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	return u.Queries.GetUserByEmail(ctx, email)
}

func (u *Users) GetUserByUsername(ctx context.Context, username string) (*db.User, error) {
	return u.Queries.GetUserByUsername(ctx, username)
}

func (u *Users) GetUserByUsernameOrEmail(ctx context.Context, value string) (*db.User, error) {
	return u.Queries.GetUserByUsernameOrEmail(ctx, value)
}

func (u *Users) GetUsersByUsernameOrFullNameAndPageInfo(ctx context.Context, value string, limit int32) (*model.Users, error) {
	var edges []*model.UsersEdge

	users, _ := u.Queries.GetUsersByUsernameOrFullNameAndPageInfo(ctx, db.GetUsersByUsernameOrFullNameAndPageInfoParams{
		Lower: value,
		Limit: limit,
	})

	for _, value := range users {
		edges = append(edges, &model.UsersEdge{Node: &db.User{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	return &model.Users{
		Edges: edges,
	}, nil
}

func (u *Users) UpdateUser(ctx context.Context, user *db.User) (*db.User, error) {
	updatedUser, _ := u.Queries.UpdateUser(ctx, db.UpdateUserParams{
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

	return updatedUser, nil
}
