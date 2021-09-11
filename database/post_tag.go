package database

import (
	"github.com/favecode/poster-core/graph/model"
	"github.com/go-pg/pg"
)

type PostTag struct {
	DB *pg.DB
}

func (p *PostTag) CreatePostTag(postTag *model.PostTag) (*model.PostTag, error) {
	_, err := p.DB.Model(postTag).Returning("*").Insert()
	return postTag, err
}
