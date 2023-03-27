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
	savedPost, err := s.Queries.GetSavedPostByUserIDAndPostID(ctx, db.GetSavedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if err != nil {
		return nil, err
	}

	return savedPost, nil
}

func (s *SavedPosts) GetSavedPostByID(ctx context.Context, id uuid.UUID) (*db.SavedPost, error) {
	savedPost, err := s.Queries.GetSavedPostByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return savedPost, nil
}

func (s *SavedPosts) GetSavedPostsByUserIDAndPageInfo(ctx context.Context, userID string, limit int32, after string) (*model.SavedPosts, error) {
	var edges []*model.SavedPostsEdge
	var endCursor string

	createdAt, _ := time.Parse(time.RFC3339, after)
	// FIXME
	UserID, _ := uuid.Parse(userID)

	savedPosts, err := s.Queries.GetSavedPostsByUserIDAndPageInfo(ctx, db.GetSavedPostsByUserIDAndPageInfoParams{
		UserID:    UserID,
		Limit:     limit,
		CreatedAt: createdAt,
	})

	totalCount, _ := s.Queries.CountSavedPostsByUserIDAndPageInfo(ctx, db.CountSavedPostsByUserIDAndPageInfoParams{
		UserID:    UserID,
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
	}, err
}

func (s *SavedPosts) DeleteSavedPostByID(ctx context.Context, id uuid.UUID) (*db.SavedPost, error) {
	// FIXME
	DeletedAt := sql.NullTime{time.Now(), true}

	savedPost, err := s.Queries.DeleteSavedPostByID(ctx, db.DeleteSavedPostByIDParams{
		ID:        id,
		DeletedAt: DeletedAt,
	})

	if err != nil {
		return nil, err
	}
	return savedPost, nil
}
