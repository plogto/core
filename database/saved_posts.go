package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type SavedPosts struct {
	DB *pg.DB
}

func (s *SavedPosts) CreateSavedPost(savedPost *model.SavedPost) (*model.SavedPost, error) {
	_, err := s.DB.Model(savedPost).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		SelectOrInsert()

	return savedPost, err
}

func (s *SavedPosts) GetSavedPostByUserIDAndPostID(userID, postID string) (*model.SavedPost, error) {
	var savedPost model.SavedPost
	err := s.DB.Model(&savedPost).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()

	return &savedPost, err
}

func (s *SavedPosts) GetSavedPostByID(id string) (*model.SavedPost, error) {
	var savedPost model.SavedPost
	err := s.DB.Model(&savedPost).Where("id = ?", id).Where("deleted_at is ?", nil).First()

	return &savedPost, err
}

func (s *SavedPosts) GetSavedPostsByUserIDAndPageInfo(userID string, limit int, after string) (*model.SavedPosts, error) {
	var savedPosts []*model.SavedPost
	var edges []*model.SavedPostsEdge
	var endCursor string

	query := s.DB.Model(&savedPosts).
		Join("INNER JOIN posts ON posts.id = saved_post.post_id").
		Join("INNER JOIN users ON users.id = posts.user_id").
		Join("INNER JOIN connections ON connections.following_id = posts.user_id").
		Where("saved_post.user_id = ?", userID).
		Where("saved_post.deleted_at is ?", nil).
		Where("posts.deleted_at is ?", nil).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.Where("users.id = ?", userID).
				WhereOr("connections.status = ?", 2).
				WhereOr("users.is_private = ?", false)
			return q, nil
		}).
		Where("connections.deleted_at is ?", nil).
		GroupExpr("saved_post.id, posts.id")

	if len(after) > 0 {
		query.Where("saved_post.created_at < ?", after)
	}

	totalCount, err :=
		query.Limit(limit).Order("saved_post.created_at DESC").SelectAndCount()

	for _, value := range savedPosts {
		edges = append(edges, &model.SavedPostsEdge{Node: &model.SavedPost{
			ID:        value.ID,
			UserID:    value.UserID,
			PostID:    value.PostID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	hasNextPage := false
	if totalCount > limit {
		hasNextPage = true
	}

	return &model.SavedPosts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (s *SavedPosts) DeleteSavedPostByID(id string) (*model.SavedPost, error) {
	DeletedAt := time.Now()
	var savedPost = &model.SavedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := s.DB.Model(savedPost).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()

	return savedPost, err
}
