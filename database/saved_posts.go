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

func (s *SavedPosts) CreatePostSave(postSave *model.SavedPost) (*model.SavedPost, error) {
	_, err := s.DB.Model(postSave).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		SelectOrInsert()

	return postSave, err
}

func (s *SavedPosts) GetSavedPostByUserIDAndPostID(userID, postID string) (*model.SavedPost, error) {
	var postSave model.SavedPost
	err := s.DB.Model(&postSave).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()

	return &postSave, err
}

func (s *SavedPosts) GetSavedPostByID(id string) (*model.SavedPost, error) {
	var postSave model.SavedPost
	err := s.DB.Model(&postSave).Where("id = ?", id).Where("deleted_at is ?", nil).First()

	return &postSave, err
}

func (s *SavedPosts) GetSavedPostsByUserIDAndPagination(userID string, limit int, after string) (*model.SavedPosts, error) {
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

	return &model.SavedPosts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (s *SavedPosts) DeletePostSaveByID(id string) (*model.SavedPost, error) {
	DeletedAt := time.Now()
	var postSave = &model.SavedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := s.DB.Model(postSave).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()

	return postSave, err
}
