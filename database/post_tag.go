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

func (p *PostTag) GetTagsOrderByCountTags(limit, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit

	query := p.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tag.tag_id").
		Join("INNER JOIN post_tag ON post_tag.tag_id = tag.id").
		Join("INNER JOIN post ON post_tag.post_id = post.id").
		Join("INNER JOIN \"user\" as u ON u.id = post.user_id").
		GroupExpr("post_tag.tag_id, tag.id").
		Where("post.deleted_at is ?", nil).
		Where("u.is_private is false").
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
