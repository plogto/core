package database

import (
	"fmt"

	"github.com/favecode/poster-core/graph/model"
	"github.com/favecode/poster-core/util"
	"github.com/go-pg/pg"
)

type Post struct {
	DB *pg.DB
}

func (p *Post) GetPostByField(field, value string) (*model.Post, error) {
	var post model.Post
	err := p.DB.Model(&post).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &post, err
}

func (p *Post) GetPostsByUserIdAndPagination(userId string, limit int, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(limit)

	totalDocs, err := query.SelectAndCount()

	return &model.Posts{
		Pagination: util.GetPatination(&util.GetPaginationParams{
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

func (p *Post) CreatePost(post *model.Post) (*model.Post, error) {
	_, err := p.DB.Model(post).Returning("*").Insert()
	return post, err
}
