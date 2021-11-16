package database

import (
	"fmt"
	"time"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg/v10"
)

type Comment struct {
	DB *pg.DB
}

func (p *Comment) CreateComment(comment *model.Comment) (*model.Comment, error) {
	_, err := p.DB.Model(comment).Returning("*").Insert()
	return comment, err
}

func (p *Comment) GetCommentByField(field, value string) (*model.Comment, error) {
	var comment model.Comment
	err := p.DB.Model(&comment).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(comment.ID) < 1 {
		return nil, nil
	}
	return &comment, err
}

func (p *Comment) GetCommentByID(id string) (*model.Comment, error) {
	return p.GetCommentByField("id", id)
}

func (p *Comment) GetCommentsByPostIdAndPagination(postId string, limit int, page int) (*model.Comments, error) {
	var comments []*model.Comment
	var offset = (page - 1) * limit

	query := p.DB.Model(&comments).
		Where("post_id = ?", postId).
		Where("parent_id is ?", nil).
		Where("deleted_at is ?", nil).
		Order("created_at DESC").
		Returning("*")

	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Comments{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Comments: comments,
	}, err
}

func (p *Comment) GetCommentsByParentIdAndPagination(parentId string, limit int, page int) (*model.Comments, error) {
	var comments []*model.Comment
	var offset = (page - 1) * limit

	query := p.DB.Model(&comments).
		Where("parent_id = ?", parentId).
		Where("deleted_at is ?", nil).
		Order("created_at DESC").
		Returning("*")

	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Comments{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		Comments: comments,
	}, err
}

func (p *Comment) DeleteCommentByID(id string) (*model.Comment, error) {
	DeletedAt := time.Now()
	var comment = &model.Comment{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(comment).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return comment, err
}
