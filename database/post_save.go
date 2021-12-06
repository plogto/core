package database

import (
	"time"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg/v10"
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

func (p *PostSave) GetPostSaveByUserIdAndPostId(userId, postId string) (*model.PostSave, error) {
	var postSave model.PostSave
	err := p.DB.Model(&postSave).Where("user_id = ?", userId).Where("post_id = ?", postId).Where("deleted_at is ?", nil).First()
	if len(postSave.ID) < 1 {
		return nil, nil
	}
	return &postSave, err
}

func (p *PostSave) GetPostSavesByUserIdAndPagination(userId string, limit, page int) (*model.PostSaves, error) {
	var postSaves []*model.PostSave
	var offset = (page - 1) * limit

	query := p.DB.Model(&postSaves).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostSaves{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		PostSaves: postSaves,
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
