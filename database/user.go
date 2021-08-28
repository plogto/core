package database

import (
	"fmt"

	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/util"
	"github.com/go-pg/pg"
)

type User struct {
	DB *pg.DB
}

func (u *User) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &user, err
}

func (u *User) GetUserByID(id string) (*model.User, error) {
	return u.GetUserByField("id", id)
}

func (u *User) GetUserByEmail(email string) (*model.User, error) {
	return u.GetUserByField("email", email)
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	return u.GetUserByField("username", username)
}

func (u *User) GetUserByUsernameOrEmail(value string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("username = ?", value).WhereOr("email = ?", value).Where("deleted_at is ?", nil).First()
	return &user, err
}

func (u *User) GetUserByUsernameOrFullnameAndPagination(value string, limit int, page int) (*model.Users, error) {
	var users []*model.User
	var offset = (page - 1) * limit

	query := u.DB.Model(&users).Where("username LIKE ?", value).WhereOr("fullname LIKE ?", value).Where("deleted_at is ?", nil)
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Users{
		Pagination: util.GetPatination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Users: users,
	}, err
}

func (u *User) CreateUser(tx *pg.Tx, user *model.User) (*model.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
