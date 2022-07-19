package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type Posts struct {
	DB *pg.DB
}

func (p *Posts) GetPostByField(field string, value string) (*model.Post, error) {
	var post model.Post
	err := p.DB.Model(&post).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	if len(post.ID) < 1 {
		return nil, nil
	}
	return &post, err
}

func (p *Posts) GetPostsByUserIDAndPagination(userID string, parentID *string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).Where("user_id = ?", userID).Where("deleted_at is ?", nil)

	if parentID != nil {
		query.Where("parent_id = ?", parentID).Order("created_at ASC")
	} else {
		query.Where("parent_id is ?", parentID).Order("created_at DESC")
	}

	query.Offset(offset).Limit(limit).Returning("*")

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

func (p *Posts) GetPostsByParentIDAndPagination(parentID string, limit, page int) (*model.Posts, error) {
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

func (p *Posts) GetPostsByTagIDAndPagination(tagID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	// TODO: fix this query
	query := p.DB.Model(&posts).
		ColumnExpr("post_tags.post_id, post_tags.tag_id").
		ColumnExpr("users.id, users.is_private").
		ColumnExpr("post.*").
		Join("INNER JOIN post_tags ON post_tags.tag_id = ?", tagID).
		Join("INNER JOIN users ON users.id = post.user_id").
		Where("post_tags.post_id = post.id").
		Where("post.deleted_at is ?", nil).
		Where("users.is_private is false").
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

func (p *Posts) GetTimelinePostsByPagination(userID string, limit, page int) (*model.Posts, error) {
	var posts []*model.Post
	var offset = (page - 1) * limit

	query := p.DB.Model(&posts).
		ColumnExpr("connections.follower_id, connections.following_id, connections.status").
		ColumnExpr("users.id, users.is_private").
		ColumnExpr("post.*").
		Join("INNER JOIN connections ON connections.follower_id = ?", userID).
		Join("INNER JOIN users ON users.id = connections.following_id").
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.Where("connections.status = ?", 2).
				WhereOr("users.is_private = ?", false)
			return q, nil
		}).
		Where("post.user_id = users.id").
		Where("post.parent_id is ?", nil).
		Where("connections.deleted_at is ?", nil).
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

func (p *Posts) GetPostByID(id string) (*model.Post, error) {
	return p.GetPostByField("id", id)
}

func (p *Posts) GetPostByURL(url string) (*model.Post, error) {
	return p.GetPostByField("url", url)
}

func (p *Posts) CountPostsByUserID(userID string) (*int, error) {
	count, err := p.DB.Model((*model.Post)(nil)).
		Where("user_id = ?", userID).
		Where("parent_id is ?", nil).
		Where("deleted_at is ?", nil).
		Count()
	return &count, err
}

func (p *Posts) CreatePost(post *model.Post) (*model.Post, error) {
	_, err := p.DB.Model(post).Returning("*").Insert()
	return post, err
}

func (p *Posts) UpdatePost(post *model.Post) (*model.Post, error) {
	_, err := p.DB.Model(post).WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return post, err
}

func (p *Posts) DeletePostByID(id string) (*model.Post, error) {
	DeletedAt := time.Now()
	var post = &model.Post{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(post).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return post, err
}
