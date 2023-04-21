package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type PostTags struct {
	Queries *db.Queries
}

func (p *PostTags) CreatePostTag(ctx context.Context, tagID, postID pgtype.UUID) (*db.PostTag, error) {
	return p.Queries.CreatePostTag(ctx, db.CreatePostTagParams{
		TagID:  tagID,
		PostID: postID,
	})
}

func (p *PostTags) GetTagsOrderByCountTags(ctx context.Context, limit int32) (*model.Tags, error) {
	var edges []*model.TagsEdge

	tags, _ := p.Queries.GetTagsOrderByCountTags(ctx, limit)

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

func (p *PostTags) CountPostTagsByTagID(ctx context.Context, tagID pgtype.UUID) (int64, error) {
	totalCount, _ := p.Queries.CountPostTagsByTagID(ctx, tagID)

	return totalCount, nil
}

func (p *PostTags) DeletePostTagsByPostID(ctx context.Context, postID pgtype.UUID) ([]*db.PostTag, error) {
	DeletedAt := time.Now()

	postTags, _ := p.Queries.DeletePostTagsByPostID(ctx, db.DeletePostTagsByPostIDParams{
		PostID:    postID,
		DeletedAt: &DeletedAt,
	})

	return postTags, nil
}
