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

func (p *PostSave) GetPostSavesByUserIDAndPagination(userID string, limit, page int) (*model.PostSaves, error) {
	var savesPosts []*model.PostSave
	var offset = (page - 1) * limit

	query := p.DB.Model(&savesPosts).Where("user_id = ?", userID).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostSaves{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		SavedPosts: savesPosts,
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
