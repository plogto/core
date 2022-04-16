package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type PostAttachment struct {
	DB *pg.DB
}

func (p *PostAttachment) CreatePostAttachment(postAttachment *model.PostAttachment) (*model.PostAttachment, error) {
	_, err := p.DB.Model(postAttachment).Returning("*").Insert()
	return postAttachment, err
}
