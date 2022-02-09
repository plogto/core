package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Post struct {
	DB *pg.DB
}

func (p *Post) GetPostByField(field string, value string) (*model.Post, error) {
	var post model.Post
	err := p.DB.Model(&post).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(post.ID) < 1 {
		return nil, nil
	}
	return &post, err
}

func (p *Post) GetPostsByUserIDAndPagination(userID string, parentID *string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("user_id = ?", userID).Where("deleted_at is ?", nil)

	if parentID != nil {
		query.Where("parent_id = ?", parentID)
	} else {
		query.Where("parent_id is ?", parentID)
	}

	query.Offset(offset).Limit(limit).Order("created_at ASC").Returning("*")

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

func (p *Post) GetPostsByParentIDAndPagination(parentID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("parent_id= ?", parentID).Where("deleted_at is ?", nil).Order("created_at ASC").Returning("*")
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

func (p *Post) GetPostsByTagIDAndPagination(tagID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	// TODO: fix this query
	query := p.DB.Model(&posts).
		ColumnExpr("post_tag.post_id, post_tag.tag_id").
		ColumnExpr("post.*").
		Join("INNER JOIN post_tag ON post_tag.tag_id = ?", tagID).
		Where("post_tag.post_id = post.id").
		Where("post.deleted_at is ?", nil).
		Order("post.created_at DESC").Returning("*")
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

func (p *Post) GetPostByID(id string) (*model.Post, error) {
	return p.GetPostByField("id", id)
}

func (p *Post) GetPostByURL(url string) (*model.Post, error) {
	return p.GetPostByField("url", url)
}

func (p *Post) CountPostsByUserID(userID string) (*int, error) {
	count, err := p.DB.Model((*model.Post)(nil)).Where("user_id = ?", userID).Where("deleted_at is ?", nil).Count()
	return &count, err
}

func (p *Post) CreatePost(post *model.Post) (*model.Post, error) {
	_, err := p.DB.Model(post).Returning("*").Insert()
	return post, err
}

func (p *Post) DeletePostByID(id string) (*model.Post, error) {
	DeletedAt := time.Now()
	var post = &model.Post{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(post).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return post, err
}
