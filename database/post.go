package database

import (
	"fmt"

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

func (p *Post) GetPostsByUserIdAndPagination(userId string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
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

func (p *Post) GetPostsByTagIdAndPagination(tagId string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	// TODO: fix this query
	query := p.DB.Model(&posts).
		ColumnExpr("post_tag.post_id, post_tag.tag_id").
		ColumnExpr("post.*").
		Join("INNER JOIN post_tag ON post_tag.tag_id = ?", tagId).
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

func (p *Post) CountPostsByUserId(userId string) (*int, error) {
	count, err := p.DB.Model((*model.Post)(nil)).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Count()
	return &count, err
}

func (p *Post) CreatePost(post *model.Post) (*model.Post, error) {
	_, err := p.DB.Model(post).Returning("*").Insert()
	return post, err
}
