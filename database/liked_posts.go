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

type LikedPosts struct {
	Queries *db.Queries
}

func (l *LikedPosts) CreateLikedPost(ctx context.Context, userID, postID uuid.UUID) (*db.LikedPost, error) {
	likedPost, _ := l.Queries.GetLikedPostByUserIDAndPostID(ctx, db.GetLikedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if likedPost != nil {
		return likedPost, nil
	}

	newLikedPost, _ := l.Queries.CreateLikedPost(ctx, db.CreateLikedPostParams{
		UserID: userID,
		PostID: postID,
	})

	return newLikedPost, nil
}

func (l *LikedPosts) GetLikedPostByID(ctx context.Context, id uuid.UUID) (*db.LikedPost, error) {
	likedPost, err := l.Queries.GetLikedPostByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return likedPost, nil
}

func (l *LikedPosts) GetLikedPostsByPostIDAndPageInfo(ctx context.Context, postID uuid.UUID, limit int32, after string) (*model.LikedPosts, error) {
	var edges []*model.LikedPostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	likedPosts, err := l.Queries.GetLikedPostsByPostIDAndPageInfo(ctx, db.GetLikedPostsByPostIDAndPageInfoParams{
		PostID:    postID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := l.Queries.CountLikedPostsByPostIDAndPageInfo(ctx, db.CountLikedPostsByPostIDAndPageInfoParams{
		PostID:    postID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	for _, value := range likedPosts {
		edges = append(edges, &model.LikedPostsEdge{Node: &db.LikedPost{
			ID:        value.ID,
			UserID:    value.UserID,
			PostID:    value.PostID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.LikedPosts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (l *LikedPosts) GetLikedPostByUserIDAndPostID(ctx context.Context, userID, postID uuid.UUID) (*db.LikedPost, error) {
	likedPost, err := l.Queries.GetLikedPostByUserIDAndPostID(ctx, db.GetLikedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if err != nil {
		return nil, err
	}

	return likedPost, nil
}

func (l *LikedPosts) GetLikedPostsByUserIDAndPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after string) (*model.LikedPosts, error) {
	var edges []*model.LikedPostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	likedPosts, err := l.Queries.GetLikedPostsByUserIDAndPageInfo(ctx, db.GetLikedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := l.Queries.CountLikedPostsByUserIDAndPageInfo(ctx, db.CountLikedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	for _, value := range likedPosts {
		edges = append(edges, &model.LikedPostsEdge{Node: &db.LikedPost{
			ID:        value.ID,
			UserID:    value.UserID,
			PostID:    value.PostID,
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

	return &model.LikedPosts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (l *LikedPosts) DeleteLikedPostByID(ctx context.Context, id uuid.UUID) (*db.LikedPost, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	likedPost, err := l.Queries.DeleteLikedPostByID(ctx, db.DeleteLikedPostByIDParams{
		ID:        id,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}
	return likedPost, nil
}
