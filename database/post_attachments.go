package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type PostAttachments struct {
	DB *pg.DB
}

func (p *PostAttachments) CreatePostAttachment(postAttachment *model.PostAttachment) (*model.PostAttachment, error) {
	_, err := p.DB.Model(postAttachment).Returning("*").Insert()
	return postAttachment, err
}
