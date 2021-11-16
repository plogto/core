package database

import (
	"time"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg/v10"
)

type CommentLike struct {
	DB *pg.DB
}

func (p *CommentLike) CreateCommentLike(commentLike *model.CommentLike) (*model.CommentLike, error) {
	_, err := p.DB.Model(commentLike).
		Where("user_id = ?user_id").
		Where("comment_id = ?comment_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return commentLike, err
}

func (p *CommentLike) GetCommentLikesByCommentIdAndPagination(commentId string, limit int, page int) (*model.CommentLikes, error) {
	var commentLikes []*model.CommentLike
	var offset = (page - 1) * limit

	query := p.DB.Model(&commentLikes).Where("comment_id = ?", commentId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.CommentLikes{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		CommentLikes: commentLikes,
	}, err
}

func (p *CommentLike) GetCommentLikeByUserIdAndCommentId(userId string, commentId string) (*model.CommentLike, error) {
	var commentLike model.CommentLike
	err := p.DB.Model(&commentLike).Where("user_id = ?", userId).Where("comment_id = ?", commentId).Where("deleted_at is ?", nil).First()
	if len(commentLike.ID) < 1 {
		return nil, nil
	}
	return &commentLike, err
}

func (p *CommentLike) DeleteCommentLikeByID(id string) (*model.CommentLike, error) {
	DeletedAt := time.Now()
	var commentLike = &model.CommentLike{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(commentLike).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return commentLike, err
}
