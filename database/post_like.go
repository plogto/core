package database

import (
	"time"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg/v10"
)

type PostLike struct {
	DB *pg.DB
}

func (p *PostLike) CreatePostLike(postLike *model.PostLike) (*model.PostLike, error) {
	_, err := p.DB.Model(postLike).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return postLike, err
}

func (p *PostLike) GetPostLikesByPostIdAndPagination(postId string, limit int, page int) (*model.PostLikes, error) {
	var postLikes []*model.PostLike
	var offset = (page - 1) * limit

	query := p.DB.Model(&postLikes).Where("post_id = ?", postId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostLikes{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		PostLikes: postLikes,
	}, err
}

func (p *PostLike) GetPostLikeByUserIdAndPostId(userId string, postId string) (*model.PostLike, error) {
	var postLike model.PostLike
	err := p.DB.Model(&postLike).Where("user_id = ?", userId).Where("post_id = ?", postId).Where("deleted_at is ?", nil).First()
	if len(postLike.ID) < 1 {
		return nil, nil
	}
	return &postLike, err
}

func (p *PostLike) DeletePostLikeByID(id string) (*model.PostLike, error) {
	DeletedAt := time.Now()
	var postLike = &model.PostLike{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(postLike).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return postLike, err
}
