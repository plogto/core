package database

import (
	"fmt"

	"github.com/favecode/note-core/config"
	"github.com/favecode/note-core/graph/model"
	"github.com/favecode/note-core/util"
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

func (p *Post) GetPostsByUserIdAndPagination(userId string, paginationParams *model.GetUserPostsByUsernameInput) (*model.Posts, error) {
	var posts []*model.Post
	var limit int = config.POSTS_PAGE_LIMIT

	if paginationParams.Limit != nil {
		limit = *paginationParams.Limit
	}

	var page int = 1
	if paginationParams.Page != nil && *paginationParams.Page > 0 {
		page = *paginationParams.Page
	}

	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Offset(offset).Limit(*paginationParams.Limit)

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
