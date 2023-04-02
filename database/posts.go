package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Posts struct {
	Queries *db.Queries
}

func (p *Posts) CreatePost(ctx context.Context, arg db.CreatePostParams) (*db.Post, error) {
	post, err := p.Queries.CreatePost(ctx, arg)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Posts) GetPostByID(ctx context.Context, id string) (*db.Post, error) {
	ID, _ := uuid.Parse(id)
	post, err := p.Queries.GetPostByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Posts) GetPostByURL(ctx context.Context, url string) (*db.Post, error) {
	post, err := p.Queries.GetPostByURL(ctx, url)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Posts) GetPostsByUserIDAndPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetPostsByUserIDAndPageInfo(ctx, db.GetPostsByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountPostsByUserIDAndPageInfo(ctx, db.CountPostsByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			ParentID:  value.ParentID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetPostsWithParentIDByUserIDAndPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetPostsWithParentIDByUserIDAndPageInfo(ctx, db.GetPostsWithParentIDByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountPostsWithParentIDByUserIDAndPageInfo(ctx, db.CountPostsWithParentIDByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			ParentID:  value.ParentID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetPostsByParentIDAndPageInfo(ctx context.Context, userID uuid.NullUUID, parentID uuid.UUID, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string
	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetPostsByParentIDAndPageInfo(ctx, db.GetPostsByParentIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		ParentID:  uuid.NullUUID{parentID, true},
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountPostsByParentIDAndPageInfo(ctx, db.CountPostsByParentIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		ParentID:  uuid.NullUUID{parentID, true},
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (p *Posts) GetPostsByTagIDAndPageInfo(ctx context.Context, tagID uuid.UUID, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetPostsByTagIDAndPageInfo(ctx, db.GetPostsByTagIDAndPageInfoParams{
		Limit:     limit,
		TagID:     tagID,
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountPostsByTagIDAndPageInfo(ctx, db.CountPostsByTagIDAndPageInfoParams{
		Limit:     limit,
		TagID:     tagID,
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetTimelinePostsByPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetTimelinePostsByPageInfo(ctx, db.GetTimelinePostsByPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountTimelinePostsByPageInfo(ctx, db.CountTimelinePostsByPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetExplorePostsByPageInfo(ctx context.Context, limit int32, after string) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	posts, err := p.Queries.GetExplorePostsByPageInfo(ctx, db.GetExplorePostsByPageInfoParams{
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := p.Queries.CountExplorePostsByPageInfo(ctx, db.CountExplorePostsByPageInfoParams{
		Limit:     limit,
		CreatedAt: createdAt,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &db.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > int64(limit) {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) CountPostsByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	count, err := p.Queries.CountPostsByUserID(ctx, userID)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *Posts) UpdatePost(ctx context.Context, post *db.Post) (*db.Post, error) {
	post, err := p.Queries.UpdatePost(ctx, db.UpdatePostParams{
		ID:      post.ID,
		Content: post.Content,
		Status:  post.Status,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Posts) DeletePostByID(ctx context.Context, id uuid.UUID) (*db.Post, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	post, err := p.Queries.DeletePostByID(ctx, db.DeletePostByIDParams{
		ID:        id,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}
