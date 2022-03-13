package database

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Tag struct {
	DB *pg.DB
}

func (t *Tag) GetTagsByTagNameAndPagination(value string, limit, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit
	value = strings.ToLower(value)

	query := t.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tag.tag_id").
		Join("INNER JOIN post_tag ON post_tag.tag_id = tag.id").
		Join("INNER JOIN post ON post_tag.post_id = post.id").
		Join("INNER JOIN \"user\" as u ON u.id = post.user_id").
		GroupExpr("post_tag.tag_id, tag.id").
		Where("lower(tag.name) LIKE lower(?)", value).
		Where("tag.deleted_at is ?", nil).
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

func (t *Tag) GetTagByField(field, value string) (*model.Tag, error) {
	var tag model.Tag
	err := t.DB.Model(&tag).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tag.tag_id").
		Join("INNER JOIN post_tag ON post_tag.tag_id = tag.id").
		Join("INNER JOIN post ON post_tag.post_id = post.id").
		Join("INNER JOIN \"user\" as u ON u.id = post.user_id").
		GroupExpr("post_tag.tag_id, tag.id").
		Where(fmt.Sprintf("lower(tag.%v) = lower(?)", field), value).
		Where("tag.deleted_at is ?", nil).
		Where("u.is_private is false").
		Returning("*").First()

	fmt.Println(tag, err)

	return &tag, err
}

func (t *Tag) GetTagByName(name string) (*model.Tag, error) {
	return t.GetTagByField("name", name)
}

func (t *Tag) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := t.DB.Model(tag).Where("name = ?name").Returning("*").SelectOrInsert()
	return tag, err
}
