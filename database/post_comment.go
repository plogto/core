package database

import (
	"fmt"

	"github.com/favecode/plog-core/graph/model"
	"github.com/favecode/plog-core/util"
	"github.com/go-pg/pg"
)

type PostComment struct {
	DB *pg.DB
}

func (p *PostComment) CreatePostComment(postComment *model.PostComment) (*model.PostComment, error) {
	_, err := p.DB.Model(postComment).Returning("*").Insert()
	return postComment, err
}

func (p *PostComment) GetPostCommentByField(field, value string) (*model.PostComment, error) {
	var postComment model.PostComment
	err := p.DB.Model(&postComment).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(postComment.ID) < 1 {
		return nil, nil
	}
	return &postComment, err
}

func (p *PostComment) GetPostCommentByID(id string) (*model.PostComment, error) {
	return p.GetPostCommentByField("id", id)
}

func (p *PostComment) GetPostCommentsByPostIdAndPagination(postId string, limit int, page int) (*model.PostComments, error) {
	var postComments []*model.PostComment
	var offset = (page - 1) * limit

	query := p.DB.Model(&postComments).
		Where("post_id = ?", postId).
		Where("parent_id is ?", nil).
		Where("deleted_at is ?", nil).
		Order("created_at DESC").
		Returning("*")

	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostComments{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		PostComments: postComments,
	}, err
}

func (p *PostComment) GetPostCommentsByParentIdAndPagination(parentId string, limit int, page int) (*model.PostComments, error) {
	var postComments []*model.PostComment
	var offset = (page - 1) * limit

	query := p.DB.Model(&postComments).
		Where("parent_id = ?", parentId).
		Where("deleted_at is ?", nil).
		Order("created_at DESC").
		Returning("*")

	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.PostComments{
		Pagination: util.GetPagination(&util.GetPaginationParams{
			Limit:     limit,
			Page:      page,
			TotalDocs: totalDocs,
		}),
		PostComments: postComments,
	}, err
}
