package database

import (
	"github.com/favecode/poster-core/graph/model"
	"github.com/go-pg/pg"
)

type Tag struct {
	DB *pg.DB
}

func (t *Tag) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := t.DB.Model(tag).Where("name = ?name").Returning("*").SelectOrInsert()
	return tag, err
}
