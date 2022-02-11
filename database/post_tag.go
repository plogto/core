package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type PostTag struct {
	DB *pg.DB
}

func (p *PostTag) CreatePostTag(postTag *model.PostTag) (*model.PostTag, error) {
	_, err := p.DB.Model(postTag).Returning("*").Insert()
	return postTag, err
}

func (p *PostTag) CountPostTagsByTagID(tagID string) (*int, error) {
	var postTags []*model.PostTag
	count, err := p.DB.Model(&postTags).
		Where("tag_id = ?", tagID).
		Where("deleted_at is ?", nil).
		Returning("*").Count()
	return &count, err
}

func (p *PostTag) GetTagsOrderByCountTags(limit, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit

	query := p.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tag.tag_id").
		Join("INNER JOIN post_tag ON post_tag.tag_id = tag.id").
		GroupExpr("post_tag.tag_id, tag.id").
		Where("tag.deleted_at is ?", nil).
		Order("count DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Tags{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Tags: tags,
	}, err
}
