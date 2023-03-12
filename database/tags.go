package database

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/plogto/core/convertor"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type Tags struct {
	Queries *db.Queries
}

func (t *Tags) CreateTag(ctx context.Context, name string) (*model.Tag, error) {
	tag, _ := t.Queries.GetTagByName(ctx, name)

	if tag != nil {
		return &model.Tag{
			ID:   tag.ID,
			Name: tag.Name,
		}, nil
	}

	newTag, _ := t.Queries.CreateTag(ctx, name)

	return &model.Tag{
		ID:   newTag.ID,
		Name: newTag.Name,
	}, nil
}

func (t *Tags) GetTagByIDs(ctx context.Context, ids []uuid.UUID) ([]*model.Tag, error) {
	tags, err := t.Queries.GetTagByIDs(ctx, ids)

	if err != nil {
		return nil, err
	}

	return convertor.DBTagsToModel(tags), nil
}
func (t *Tags) GetTagByName(ctx context.Context, name string) (*model.Tag, error) {
	tag, err := t.Queries.GetTagByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return &model.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (t *Tags) GetTagsByTagNameAndPageInfo(ctx context.Context, name string, limit int) (*model.Tags, error) {
	var edges []*model.TagsEdge

	name = strings.ToLower(name)

	Limit := int32(limit)
	tags, err := t.Queries.GetTagsByTagNameAndPageInfo(ctx, db.GetTagsByTagNameAndPageInfoParams{
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
	}, err
}
