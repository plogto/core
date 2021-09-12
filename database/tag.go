package database

import (
	"fmt"
	"strings"

	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/util"
	"github.com/go-pg/pg"
)

type Tag struct {
	DB *pg.DB
}

func (t *Tag) GetTagsByTagNameAndPagination(value string, limit int, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit
	value = strings.ToLower(value)

	query := t.DB.Model(&tags).Where("lower(name) LIKE ?", value).Where("deleted_at is ?", nil)
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Tags{
		Pagination: util.GetPatination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Tags: tags,
	}, err
}

func (t *Tag) GetTagByField(field, value string) (*model.Tag, error) {
	var tag model.Tag
	err := t.DB.Model(&tag).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &tag, err
}

func (t *Tag) GetTagByName(name string) (*model.Tag, error) {
	return t.GetTagByField("name", name)
}

func (t *Tag) CreateTag(tag *model.Tag) (*model.Tag, error) {
	_, err := t.DB.Model(tag).Where("name = ?name").Returning("*").SelectOrInsert()
	return tag, err
}
