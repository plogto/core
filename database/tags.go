package database

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Tags struct {
	DB *pg.DB
}

func (t *Tags) GetTagsByTagNameAndPagination(value string, limit, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit
	value = strings.ToLower(value)

	query := t.DB.Model(&tags).
		ColumnExpr("tags.*, count(tags.id)").
		ColumnExpr("post_tags.tag_id").
		Join("INNER JOIN post_tags ON post_tags.tag_id = tags.id").
		Join("INNER JOIN posts ON post_tags.post_id = posts.id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		GroupExpr("post_tags.tag_id, tags.id").
		Where("lower(tags.name) LIKE lower(?)", strings.ToLower(value)).
		Where("posts.deleted_at is ?", nil).
		Where("users.is_private is false").
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

func (t *Tags) GetTagByField(field, value string) (*model.Tag, error) {
	var tag model.Tag
	err := t.DB.Model(&tag).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tags.tag_id").
		Join("INNER JOIN post_tags ON post_tags.tag_id = tag.id").
		Join("INNER JOIN posts ON post_tags.post_id = posts.id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		GroupExpr("post_tags.tag_id, tag.id").
		Where(fmt.Sprintf("lower(tag.%v) = lower(?)", field), value).
		Where("posts.deleted_at is ?", nil).
		Where("users.is_private is false").
		Returning("*").First()

	return &tag, err
}

func (t *Tags) GetTagByName(name string) (*model.Tag, error) {
	return t.GetTagByField("name", name)
}

func (t *Tags) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := t.DB.Model(tag).Where("name = ?name").Group("id").Returning("*").SelectOrInsert()
	return tag, err
}
