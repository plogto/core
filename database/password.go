package database

import (
	"fmt"

	"github.com/favecode/plog-core/graph/model"
	"github.com/go-pg/pg/v10"
)

type Password struct {
	DB *pg.DB
}

func (p *Password) GetPasswordByField(field, value string) (*model.Password, error) {
	var password model.Password
	err := p.DB.Model(&password).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &password, err
}

func (p *Password) GetPasswordByUserID(id string) (*model.Password, error) {
	return p.GetPasswordByField("user_id", id)
}

func (p *Password) AddPassword(password *model.Password) (*model.Password, error) {
	_, err := p.DB.Model(password).Returning("*").Insert()
	return password, err
}

func (p *Password) UpdatePassword(password *model.Password) (*model.Password, error) {
	_, err := p.DB.Model(password).WherePK().Returning("*").Update()
	return password, err
}
