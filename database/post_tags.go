package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
)

type PostTags struct {
	DB *pg.DB
}

func (p *PostTags) CreatePostTag(postTag *model.PostTag) (*model.PostTag, error) {
	_, err := p.DB.Model(postTag).Returning("*").Insert()
	return postTag, err
}

func (p *PostTags) GetTagsOrderByCountTags(limit int) (*model.Tags, error) {
	var tags []*model.Tag
	var edges []*model.TagsEdge

	err := p.DB.Model(&tags).
		ColumnExpr("tag.*, count(tag.id)").
		Join("INNER JOIN post_tags ON post_tags.tag_id = tag.id").
		Join("INNER JOIN posts ON post_tags.post_id = posts.id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		GroupExpr("post_tags.tag_id, tag.id").
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
