package database

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

type Tags struct {
	Queries *db.Queries
}

func (t *Tags) CreateTag(ctx context.Context, name string) (*model.Tag, error) {
	tag := util.HandleDBResponseWithoutError(t.Queries.GetTagByName(ctx, name))

	if validation.IsTagExists(tag) {
		return &model.Tag{
			ID:   tag.ID,
			Name: tag.Name,
		}, nil
	}

	newTag, _ := t.Queries.CreateTag(ctx, name)

	return util.HandleDBResponse(&model.Tag{
		ID:   newTag.ID,
		Name: newTag.Name,
	}, nil)
}

func (t *Tags) GetTagByIDs(ctx context.Context, ids []uuid.UUID) ([]*model.Tag, error) {
	tags, _ := t.Queries.GetTagByIDs(ctx, ids)

	return convertor.DBTagsToModel(tags), nil
}
func (t *Tags) GetTagByName(ctx context.Context, name string) (*model.Tag, error) {
	tag, err := t.Queries.GetTagByName(ctx, name)

	return util.HandleDBResponse(&model.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}, err)
}

func (t *Tags) GetTagsByTagNameAndPageInfo(ctx context.Context, name string, limit int) (*model.Tags, error) {
	var edges []*model.TagsEdge

	name = strings.ToLower(name)

	Limit := int32(limit)
	tags, _ := t.Queries.GetTagsByTagNameAndPageInfo(ctx, db.GetTagsByTagNameAndPageInfoParams{
		Name:  name,
		Limit: Limit,
	})

	for _, value := range tags {
		edges = append(edges, &model.TagsEdge{Node: &model.Tag{
			ID:        value.ID,
			Count:     value.Count,
			CreatedAt: value.CreatedAt,
		}})
	}

	return &model.Tags{
		Edges: edges,
	}, nil
}
