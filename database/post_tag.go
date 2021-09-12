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

func (p *PostTag) CountPostTagsByTagId(tagId string) (*int, error) {
	var postTags []*model.PostTag
	count, err := p.DB.Model(&postTags).
		Where("tag_id = ?", tagId).
		Where("deleted_at is ?", nil).
		Returning("*").Count()
	return &count, err
}
