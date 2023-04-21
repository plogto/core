package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Posts struct {
	Queries *db.Queries
}

func (p *Posts) CreatePost(ctx context.Context, arg db.CreatePostParams) (*db.Post, error) {
	return p.Queries.CreatePost(ctx, arg)
}

func (p *Posts) GetPostByURL(ctx context.Context, url string) (*db.Post, error) {
	return p.Queries.GetPostByURL(ctx, url)
}

func (p *Posts) GetPostsByUserIDAndPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsByUserIDAndPageInfo(ctx, db.GetPostsByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsByUserIDAndPageInfo(ctx, db.CountPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetPostsWithAttachmentByUserIDAndPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsWithAttachmentByUserIDAndPageInfo(ctx, db.GetPostsWithAttachmentByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsWithAttachmentByUserIDAndPageInfo(ctx, db.CountPostsWithAttachmentByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetPostsWithParentIDByUserIDAndPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsWithParentIDByUserIDAndPageInfo(ctx, db.GetPostsWithParentIDByUserIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsWithParentIDByUserIDAndPageInfo(ctx, db.CountPostsWithParentIDByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetPostsByUserIDAndParentIDAndPageInfo(ctx context.Context, userID pgtype.UUID, parentID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsByUserIDAndParentIDAndPageInfo(ctx, db.GetPostsByUserIDAndParentIDAndPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		ParentID:  parentID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsByUserIDAndParentIDAndPageInfo(ctx, db.CountPostsByUserIDAndParentIDAndPageInfoParams{
		UserID:    userID,
		ParentID:  parentID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
	}, nil
}

func (p *Posts) GetPostsByParentIDAndPageInfo(ctx context.Context, parentID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsByParentIDAndPageInfo(ctx, db.GetPostsByParentIDAndPageInfoParams{
		Limit:     limit,
		ParentID:  parentID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsByParentIDAndPageInfo(ctx, db.CountPostsByParentIDAndPageInfoParams{
		ParentID:  parentID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
	}, nil
}

func (p *Posts) GetPostsByTagIDAndPageInfo(ctx context.Context, tagID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetPostsByTagIDAndPageInfo(ctx, db.GetPostsByTagIDAndPageInfoParams{
		Limit:     limit,
		TagID:     tagID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountPostsByTagIDAndPageInfo(ctx, db.CountPostsByTagIDAndPageInfoParams{
		TagID:     tagID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetTimelinePostsByPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetTimelinePostsByPageInfo(ctx, db.GetTimelinePostsByPageInfoParams{
		Limit:     limit,
		UserID:    userID,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountTimelinePostsByPageInfo(ctx, db.CountTimelinePostsByPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
	})

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetExplorePostsByPageInfo(ctx context.Context, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetExplorePostsByPageInfo(ctx, db.GetExplorePostsByPageInfoParams{
		Limit:     limit,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountExplorePostsByPageInfo(ctx, after)

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) GetExplorePostsWithAttachmentByPageInfo(ctx context.Context, limit int32, after time.Time) (*model.Posts, error) {
	var edges []*model.PostsEdge
	var endCursor string

	posts, _ := p.Queries.GetExplorePostsWithAttachmentByPageInfo(ctx, db.GetExplorePostsWithAttachmentByPageInfoParams{
		Limit:     limit,
		CreatedAt: after,
	})

	totalCount, _ := p.Queries.CountExplorePostsWithAttachmentByPageInfo(ctx, after)

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (p *Posts) CountPostsByUserID(ctx context.Context, userID pgtype.UUID) (int64, error) {
	count, _ := p.Queries.CountPostsByUserID(ctx, userID)

	return count, nil
}

func (p *Posts) UpdatePost(ctx context.Context, post *db.Post) (*db.Post, error) {
	return p.Queries.UpdatePost(ctx, db.UpdatePostParams{
		ID:      post.ID,
		Content: post.Content,
		Status:  post.Status,
	})
}

func (p *Posts) DeletePostByID(ctx context.Context, id pgtype.UUID) (*db.Post, error) {
	DeletedAt := time.Now()

	post, _ := p.Queries.DeletePostByID(ctx, db.DeletePostByIDParams{
		ID:        id,
		DeletedAt: &DeletedAt,
	})

	return post, nil
}
