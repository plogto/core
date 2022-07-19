package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type Passwords struct {
	DB *pg.DB
}

func (p *Passwords) GetPasswordByField(field, value string) (*model.Password, error) {
	var password model.Password
	err := p.DB.Model(&password).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &password, err
}

func (p *Passwords) GetPasswordByUserID(id string) (*model.Password, error) {
	return p.GetPasswordByField("user_id", id)
}

func (p *Passwords) AddPassword(password *model.Password) (*model.Password, error) {
	_, err := p.DB.Model(password).Returning("*").Insert()
	return password, err
}

func (p *Passwords) UpdatePassword(password *model.Password) (*model.Password, error) {
	_, err := p.DB.Model(password).WherePK().Returning("*").Update()
	return password, err
}
