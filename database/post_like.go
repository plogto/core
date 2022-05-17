package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
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

func (p *PostLike) GetPostLikesByPostIDAndPagination(postID string, limit, page int) (*model.PostLikes, error) {
	var likedPosts []*model.PostLike
	var offset = (page - 1) * limit

	query := p.DB.Model(&likedPosts).Where("post_id = ?", postID).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostLikes{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		LikedPosts: likedPosts,
	}, err
}

func (p *PostLike) GetPostLikeByUserIDAndPostID(userID, postID string) (*model.PostLike, error) {
	var postLike model.PostLike
	err := p.DB.Model(&postLike).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()
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
	_, err := p.DB.Model(postLike).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return postLike, err
}
