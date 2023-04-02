package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

type PostTags struct {
	Queries *db.Queries
}

func (p *PostTags) CreatePostTag(ctx context.Context, tagID, postID string) (*db.PostTag, error) {
	TagID, _ := uuid.Parse(tagID)
	PostID, _ := uuid.Parse(postID)

	postTag, _ := p.Queries.CreatePostTag(ctx, db.CreatePostTagParams{
		TagID:  TagID,
		PostID: PostID,
	})

	return postTag, nil
}

func (p *PostTags) GetTagsOrderByCountTags(ctx context.Context, limit int) (*model.Tags, error) {
	var edges []*model.TagsEdge

	Limit := int32(limit)
	tags, _ := p.Queries.GetTagsOrderByCountTags(ctx, Limit)

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

func (p *PostTags) CountPostTagsByTagID(ctx context.Context, tagID string) (int64, error) {
	TagID, _ := uuid.Parse(tagID)
	totalCount, _ := p.Queries.CountPostTagsByTagID(ctx, TagID)

	return totalCount, nil
}

func (p *PostTags) DeletePostTagsByPostID(ctx context.Context, postID string) ([]*db.PostTag, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	PostID, _ := uuid.Parse(postID)
	postTags, _ := p.Queries.DeletePostTagsByPostID(ctx, db.DeletePostTagsByPostIDParams{
		PostID:    PostID,
		DeletedAt: DeletedAt,
	})

	return postTags, nil
}
