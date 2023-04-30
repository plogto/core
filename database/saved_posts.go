package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
	"github.com/plogto/core/validation"
)

type SavedPosts struct {
	Queries *db.Queries
}

func (s *SavedPosts) CreateSavedPost(ctx context.Context, userID, postID pgtype.UUID) (*db.SavedPost, error) {
	savedPost, _ := s.Queries.GetSavedPostByUserIDAndPostID(ctx, db.GetSavedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})

	if validation.IsSavedPostExists(savedPost) {
		return savedPost, nil
	}

	newSavedPost, _ := s.Queries.CreateSavedPost(ctx, db.CreateSavedPostParams{
		UserID: userID,
		PostID: postID,
	})

	return newSavedPost, nil
}

func (s *SavedPosts) GetSavedPostByUserIDAndPostID(ctx context.Context, userID, postID pgtype.UUID) (*db.SavedPost, error) {
	return s.Queries.GetSavedPostByUserIDAndPostID(ctx, db.GetSavedPostByUserIDAndPostIDParams{
		UserID: userID,
		PostID: postID,
	})
}

func (s *SavedPosts) GetSavedPostByID(ctx context.Context, id pgtype.UUID) (*db.SavedPost, error) {
	savedPost, err := s.Queries.GetSavedPostByID(ctx, id)
	return savedPost, err
}

func (s *SavedPosts) GetSavedPostsByUserIDAndPageInfo(ctx context.Context, userID pgtype.UUID, limit int32, after time.Time) (*model.SavedPosts, error) {
	var edges []*model.SavedPostsEdge
	var endCursor string

	savedPosts, _ := s.Queries.GetSavedPostsByUserIDAndPageInfo(ctx, db.GetSavedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		Limit:     limit,
		CreatedAt: after,
	})

	totalCount, _ := s.Queries.CountSavedPostsByUserIDAndPageInfo(ctx, db.CountSavedPostsByUserIDAndPageInfoParams{
		UserID:    userID,
		CreatedAt: after,
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
			HasNextPage: hasNextPage,
		},
	}, nil
}

func (s *SavedPosts) DeleteSavedPostByID(ctx context.Context, id pgtype.UUID) (*db.SavedPost, error) {
	DeletedAt := time.Now()

	savedPost, _ := s.Queries.DeleteSavedPostByID(ctx, db.DeleteSavedPostByIDParams{
		ID:        id,
		DeletedAt: &DeletedAt,
	})

	return savedPost, nil
}
