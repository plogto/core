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

func (p *SavedPosts) CreatePostSave(postSave *model.SavedPost) (*model.SavedPost, error) {
	_, err := p.DB.Model(postSave).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return postSave, err
}

func (p *SavedPosts) GetPostSaveByUserIDAndPostID(userID, postID string) (*model.SavedPost, error) {
	var postSave model.SavedPost
	err := p.DB.Model(&postSave).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()
	if len(postSave.ID) < 1 {
		return nil, nil
	}
	return &postSave, err
}
func (p *SavedPosts) GetSavedPostsByUserIDAndPagination(userID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	// TODO: fix this query
	query := p.DB.Model(&posts).
		ColumnExpr("saved_posts.post_id, saved_posts.user_id").
		ColumnExpr("users.id").
		ColumnExpr("post.*").
		Join("INNER JOIN saved_posts ON saved_posts.user_id = ?", userID).
		Join("INNER JOIN users ON users.id = post.user_id").
		Where("saved_posts.post_id = post.id").
		Where("saved_posts.deleted_at is ?", nil).
		Where("post.deleted_at is ?", nil).
		Order("saved_posts.created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()
	return &model.Posts{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Posts: posts,
	}, err
}

func (p *SavedPosts) DeletePostSaveByID(id string) (*model.SavedPost, error) {
	DeletedAt := time.Now()
	var postSave = &model.SavedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(postSave).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return postSave, err
}
