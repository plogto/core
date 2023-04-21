package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type LikedPosts struct {
	Queries *db.Queries
}

func (l *LikedPosts) CreateLikedPost(ctx context.Context, userID, postID pgtype.UUID) (*db.LikedPost, error) {
	likedPost, _ := l.Queries.GetLikedPostByUserIDAndPostID(ctx, db.GetLikedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if likedPost != nil {
		return likedPost, nil
	}

	return l.Queries.CreateLikedPost(ctx, db.CreateLikedPostParams{
		UserID: userID,
		PostID: postID,
	})
}

func (l *LikedPosts) GetLikedPostByID(ctx context.Context, id pgtype.UUID) (*db.LikedPost, error) {
	return l.Queries.GetLikedPostByID(ctx, id)
}

func (l *LikedPosts) GetLikedPostsByPostIDAndPageInfo(ctx context.Context, postID pgtype.UUID, limit int32, after time.Time) (*model.LikedPosts, error) {
	var edges []*model.LikedPostsEdge
	var endCursor string

	likedPosts, _ := l.Queries.GetLikedPostsByPostIDAndPageInfo(ctx, db.GetLikedPostsByPostIDAndPageInfoParams{
		PostID:    postID,
		Limit:     limit,
		CreatedAt: after,
	})

	totalCount, _ := l.Queries.CountLikedPostsByPostIDAndPageInfo(ctx, db.CountLikedPostsByPostIDAndPageInfoParams{
		PostID:    postID,
		CreatedAt: after,
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
	}, nil
}

func (l *LikedPosts) GetLikedPostByUserIDAndPostID(ctx context.Context, userID, postID pgtype.UUID) (*db.LikedPost, error) {
	return l.Queries.GetLikedPostByUserIDAndPostID(ctx, db.GetLikedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

}

func (l *LikedPosts) GetLikedPostsByUserIDAndPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.LikedPosts, error) {
	var edges []*model.LikedPostsEdge
	var endCursor string

	likedPosts, _ := l.Queries.GetLikedPostsByUserIDAndPageInfo(ctx, db.GetLikedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: after,
	})

	totalCount, _ := l.Queries.CountLikedPostsByUserIDAndPageInfo(ctx, db.CountLikedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (l *LikedPosts) DeleteLikedPostByID(ctx context.Context, id pgtype.UUID) (*db.LikedPost, error) {
	DeletedAt := time.Now()

	likedPost, _ := l.Queries.DeleteLikedPostByID(ctx, db.DeleteLikedPostByIDParams{
		ID:        id,
		DeletedAt: &DeletedAt,
	})

	return likedPost, nil
}
