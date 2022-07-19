package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type LikedPosts struct {
	DB *pg.DB
}

func (p *LikedPosts) CreatePostLike(postLike *model.LikedPost) (*model.LikedPost, error) {
	_, err := p.DB.Model(postLike).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return postLike, err
}

func (p *LikedPosts) GetLikedPostsByPostIDAndPagination(postID string, limit, page int) (*model.LikedPosts, error) {
	var likedPosts []*model.LikedPost
	var offset = (page - 1) * limit

	query := p.DB.Model(&likedPosts).Where("post_id = ?", postID).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.LikedPosts{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		LikedPosts: likedPosts,
	}, err
}

func (p *LikedPosts) GetPostLikeByUserIDAndPostID(userID, postID string) (*model.LikedPost, error) {
	var postLike model.LikedPost
	err := p.DB.Model(&postLike).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()
	if len(postLike.ID) < 1 {
		return nil, nil
	}
	return &postLike, err
}

func (p *LikedPosts) DeletePostLikeByID(id string) (*model.LikedPost, error) {
	DeletedAt := time.Now()
	var postLike = &model.LikedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(postLike).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return postLike, err
}
