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
	err := p.DB.Model(&post).
		Where(fmt.Sprintf("%v = ?", field), value).
		Where("deleted_at is ?", nil).
		First()

	return &post, err
}

func (p *Posts) GetPostsByUserIDAndPageInfo(userID string, parentID *string, limit int, after string) (*model.Posts, error) {
	var posts []*model.Post
	var edges []*model.PostsEdge

	query := p.DB.Model(&posts).
		Where("user_id = ?", userID).
		Where("deleted_at is ?", nil).
		Order("created_at DESC")

	if parentID != nil {
		query.Where("parent_id = ?", parentID)
	} else {
		query.Where("parent_id is ?", parentID)
	}

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
			ID:        value.ID,
			ParentID:  value.ParentID,
			CreatedAt: value.CreatedAt,
		}})
	}

	endCursor := util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)

	hasNextPage := false
	if totalCount > limit {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetPostsByParentIDAndPageInfo(parentID string, limit int, after string) (*model.Posts, error) {
	var posts []*model.Post
	var edges []*model.PostsEdge
	var endCursor string

	query := p.DB.Model(&posts).
		Where("parent_id= ?", parentID).
		Where("deleted_at is ?", nil).
		Order("created_at DESC")

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.Posts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (p *Posts) GetPostsByTagIDAndPageInfo(tagID string, limit int, after string) (*model.Posts, error) {
	var posts []*model.Post
	var edges []*model.PostsEdge
	var endCursor string

	// TODO: extend this query for following users
	query := p.DB.Model(&posts).
		Join("INNER JOIN post_tags ON post_tags.tag_id = ?", tagID).
		Join("INNER JOIN users ON users.id = post.user_id").
		Where("post_tags.post_id = post.id").
		Where("post.deleted_at is ?", nil).
		Where("users.is_private is false").
		Order("post.created_at DESC")

	if len(after) > 0 {
		query.Where("post.created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).SelectAndCount()

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.Posts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err
}

func (p *Posts) GetTimelinePostsByPageInfo(userID string, limit int, after string) (*model.Posts, error) {
	var followingPosts []*model.Post
	var userPosts []*model.Post
	var posts []*model.Post
	var edges []*model.PostsEdge

	followingPostsQuery := p.DB.Model(&followingPosts).
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
		Order("post.created_at DESC")

	userPostsQuery := p.DB.Model(&userPosts).
		Where("post.user_id = ?", userID).
		Where("post.parent_id is ?", nil).
		Where("post.deleted_at is ?", nil)

	userPostsQuery.Union(followingPostsQuery)

	query := p.DB.Model(&posts).With("posts", userPostsQuery)

	if len(after) > 0 {
		query.Where("post.created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).Order("post.created_at DESC").SelectAndCount()

	for _, value := range posts {
		edges = append(edges, &model.PostsEdge{Node: &model.Post{
			ID:        value.ID,
			CreatedAt: value.CreatedAt,
		}})
	}

	endCursor := util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)

	hasNextPage := false
	if totalCount > limit {
		hasNextPage = true
	}

	return &model.Posts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor:   endCursor,
			HasNextPage: &hasNextPage,
		},
	}, err
}

func (p *Posts) GetPostByID(id string) (*model.Post, error) {
	return p.GetPostByField("id", id)
}

func (p *Posts) GetPostByURL(url string) (*model.Post, error) {
	return p.GetPostByField("url", url)
}

func (p *Posts) CountPostsByUserID(userID string) (int, error) {
	return p.DB.Model((*model.Post)(nil)).
		Where("user_id = ?", userID).
		Where("parent_id is ?", nil).
		Where("deleted_at is ?", nil).
		Count()
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
