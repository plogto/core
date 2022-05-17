package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type PostSave struct {
	DB *pg.DB
}

func (p *PostSave) CreatePostSave(postSave *model.PostSave) (*model.PostSave, error) {
	_, err := p.DB.Model(postSave).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return postSave, err
}

func (p *PostSave) GetPostSaveByUserIDAndPostID(userID, postID string) (*model.PostSave, error) {
	var postSave model.PostSave
	err := p.DB.Model(&postSave).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()
	if len(postSave.ID) < 1 {
		return nil, nil
	}
	return &postSave, err
}
func (p *PostSave) GetSavedPostsByUserIDAndPagination(userID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	// TODO: fix this query
	query := p.DB.Model(&posts).
		ColumnExpr("post_save.post_id, post_save.user_id").
		ColumnExpr("u.id").
		ColumnExpr("post.*").
		Join("INNER JOIN post_save ON post_save.user_id = ?", userID).
		Join("INNER JOIN \"user\" as u ON u.id = post.user_id").
		Where("post_save.post_id = post.id").
		Where("post_save.deleted_at is ?", nil).
		Where("post.deleted_at is ?", nil).
		Order("post_save.created_at DESC").Returning("*")
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

func (p *PostSave) DeletePostSaveByID(id string) (*model.PostSave, error) {
	DeletedAt := time.Now()
	var postSave = &model.PostSave{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(postSave).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return postSave, err
}
