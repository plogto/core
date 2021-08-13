package database

import (
	"fmt"

	"github.com/favecode/note-core/graph/model"
	"github.com/go-pg/pg"
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

func (p *Password) AddPassword(tx *pg.Tx, password *model.Password) (*model.Password, error) {
	_, err := tx.Model(password).Returning("*").Insert()
	return password, err
}
