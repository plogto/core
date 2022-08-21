package database

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type Tags struct {
	DB *pg.DB
}

func (t *Tags) GetTagsByTagNameAndPageInfo(value string, limit int) (*model.Tags, error) {
	var tags []*model.Tag
	var edges []*model.TagsEdge

	value = strings.ToLower(value)

	err := t.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tags.tag_id").
		Join("INNER JOIN post_tags ON post_tags.tag_id = tag.id").
		Join("INNER JOIN posts ON post_tags.post_id = posts.id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		GroupExpr("post_tags.tag_id, tag.id").
		Where("lower(tag.name) LIKE lower(?)", strings.ToLower(value)).
		Where("posts.deleted_at is ?", nil).
		Where("users.is_private is false").
		Order("count DESC").
		Limit(limit).
		Select()

	for _, value := range tags {
		edges = append(edges, &model.TagsEdge{Node: &model.Tag{
			ID:        value.ID,
			Count:     value.Count,
			CreatedAt: value.CreatedAt,
		}})
	}

	return &model.Tags{
		Edges: edges,
	}, err
}

func (t *Tags) GetTagByField(field, value string) (*model.Tag, error) {
	var tag model.Tag
	err := t.DB.Model(&tag).
		Where(fmt.Sprintf("tag.%v = ?", field), value).First()

	return &tag, err
}

func (t *Tags) GetTagByName(value string) (*model.Tag, error) {
	var tag model.Tag
	err := t.DB.Model(&tag).
		Where("lower(tag.name) = lower(?)", value).First()

	return &tag, err
}

func (t *Tags) GetTagByID(id string) (*model.Tag, error) {
	return t.GetTagByField("id", id)
}

func (t *Tags) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := t.DB.Model(tag).Where("name = ?name").Group("id").Returning("*").SelectOrInsert()
	return tag, err
}
