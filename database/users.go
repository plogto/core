package database

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Users struct {
	DB *pg.DB
}

func (u *Users) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	value = strings.ToLower(value)
	err := u.DB.Model(&user).Where(fmt.Sprintf("lower(%v) = lower(?)", field), value).Where("deleted_at is ?", nil).First()

	return &user, err
}

func (u *Users) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).Where("deleted_at is ?", nil).First()

	return &user, err
}

func (u *Users) GetUserByEmail(email string) (*model.User, error) {
	return u.GetUserByField("email", email)
}

func (u *Users) GetUserByUsername(username string) (*model.User, error) {
	return u.GetUserByField("username", username)
}

func (u *Users) GetUserByUsernameOrEmail(value string) (*model.User, error) {
	var user model.User
	value = strings.ToLower(value)
	err := u.DB.Model(&user).Where("lower(username) = lower(?)", value).WhereOr("lower(email) = lower(?)", value).Where("deleted_at is ?", nil).First()

	return &user, err
}

func (u *Users) GetUsersByUsernameOrFullNameAndPagination(value string, limit, page int) (*model.Users, error) {
	var users []*model.User
	var offset = (page - 1) * limit
	value = strings.ToLower(value)

	query := u.DB.Model(&users).Where("lower(username) LIKE lower(?)", value).WhereOr("lower(full_name) LIKE lower(?)", value).Where("deleted_at is ?", nil)
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Users{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Users: users,
	}, err
}

func (u *Users) CreateUser(user *model.User) (*model.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()

	return user, err
}

func (u *Users) UpdateUser(user *model.User) (*model.User, error) {
	_, err := u.DB.Model(user).WherePK().Where("deleted_at is ?", nil).Returning("*").Update()

	return user, err
}
