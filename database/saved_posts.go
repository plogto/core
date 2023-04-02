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

type SavedPosts struct {
	Queries *db.Queries
}

func (s *SavedPosts) CreateSavedPost(ctx context.Context, userID, postID uuid.UUID) (*db.SavedPost, error) {
	savedPost, _ := s.Queries.GetSavedPostByUserIDAndPostID(ctx, db.GetSavedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if savedPost != nil {
		return savedPost, nil
	}

	newSavedPost, _ := s.Queries.CreateSavedPost(ctx, db.CreateSavedPostParams{
		UserID: userID,
		PostID: postID,
	})

	return newSavedPost, nil
}

func (s *SavedPosts) GetSavedPostByUserIDAndPostID(ctx context.Context, userID, postID uuid.UUID) (*db.SavedPost, error) {
	savedPost, _ := s.Queries.GetSavedPostByUserIDAndPostID(ctx, db.GetSavedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	return savedPost, nil
}

func (s *SavedPosts) GetSavedPostByID(ctx context.Context, id uuid.UUID) (*db.SavedPost, error) {
	savedPost, _ := s.Queries.GetSavedPostByID(ctx, id)

	return savedPost, nil
}

func (s *SavedPosts) GetSavedPostsByUserIDAndPageInfo(ctx context.Context, userID uuid.UUID, limit int32, after string) (*model.SavedPosts, error) {
	var edges []*model.SavedPostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)

	savedPosts, _ := s.Queries.GetSavedPostsByUserIDAndPageInfo(ctx, db.GetSavedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := s.Queries.CountSavedPostsByUserIDAndPageInfo(ctx, db.CountSavedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	for _, value := range savedPosts {
		edges = append(edges, &model.SavedPostsEdge{Node: &db.SavedPost{
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

	return &model.SavedPosts{
		TotalCount: totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, nil
}

func (s *SavedPosts) DeleteSavedPostByID(ctx context.Context, id uuid.UUID) (*db.SavedPost, error) {
	DeletedAt := sql.NullTime{time.Now(), true}

	savedPost, _ := s.Queries.DeleteSavedPostByID(ctx, db.DeleteSavedPostByIDParams{
		ID:        id,
		DeletedAt: DeletedAt,
	})

	return savedPost, nil
}
