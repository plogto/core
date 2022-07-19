package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type PostTags struct {
	DB *pg.DB
}

func (p *PostTags) CreatePostTag(postTag *model.PostTag) (*model.PostTag, error) {
	_, err := p.DB.Model(postTag).Returning("*").Insert()
	return postTag, err
}

func (p *PostTags) GetTagsOrderByCountTags(limit, page int) (*model.Tags, error) {
	var tags []*model.Tag
	var offset = (page - 1) * limit

	query := p.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		ColumnExpr("post_tags.tag_id").
		Join("INNER JOIN post_tags ON post_tags.tag_id = tag.id").
		Join("INNER JOIN posts ON post_tags.post_id = posts.id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		GroupExpr("post_tags.tag_id, tag.id").
		Where("posts.deleted_at is ?", nil).
		Where("users.is_private is false").
		Order("count DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	err := query.Select()

	fmt.Println(err)

	return &model.Tags{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: 0,
		}),
		Tags: tags,
	}, err
}
