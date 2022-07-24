package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type SavedPosts struct {
	DB *pg.DB
}

func (p *SavedPosts) CreatePostSave(postSave *model.SavedPost) (*model.SavedPost, error) {
	_, err := p.DB.Model(postSave).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		SelectOrInsert()

	return postSave, err
}

func (p *SavedPosts) GetPostSaveByUserIDAndPostID(userID, postID string) (*model.SavedPost, error) {
	var postSave model.SavedPost
	err := p.DB.Model(&postSave).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()

	return &postSave, err
}

// TODO: return array of SavedPost instead of Post
func (p *SavedPosts) GetSavedPostsByUserIDAndPagination(userID string, limit int, after string) (*model.Posts, error) {
	var followingPosts []*model.Post
	var userPosts []*model.Post
	var posts []*model.Post
	var edges []*model.PostsEdge
	var endCursor string

	followingPostsQuery := p.DB.Model(&followingPosts).
		Join("INNER JOIN saved_posts ON saved_posts.user_id = ?", userID).
		Join("INNER JOIN users ON users.id = post.user_id").
		Join("INNER JOIN connections ON connections.following_id = post.user_id").
		Where("saved_posts.post_id = post.id").
		Where("saved_posts.deleted_at is ?", nil).
		Where("post.deleted_at is ?", nil).
		Where("connections.follower_id = saved_posts.user_id").
		Where("connections.status = ?", 2).
		Where("connections.deleted_at is ?", nil)

	userPostsQuery := p.DB.Model(&userPosts).
		Join("INNER JOIN saved_posts ON saved_posts.user_id = ?", userID).
		Join("INNER JOIN users ON users.id = post.user_id").
		Where("saved_posts.post_id = post.id").
		Where("saved_posts.deleted_at is ?", nil).
		Where("post.deleted_at is ?", nil).
		WhereGroup(func(q *pg.Query) (*pg.Query, error) {
			q = q.Where("users.id = ?", userID).
				WhereOr("users.is_private = ?", false)
			return q, nil
		})

	userPostsQuery.Union(followingPostsQuery)

	query := p.DB.Model(&posts).With("posts", userPostsQuery)

	if len(after) > 0 {
		query.Where("saved_posts.created_at < ?", after)
	}

	totalCount, err := query.Limit(limit).Order("saved_posts.created_at DESC").SelectAndCount()

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

func (p *SavedPosts) DeletePostSaveByID(id string) (*model.SavedPost, error) {
	DeletedAt := time.Now()
	var postSave = &model.SavedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := p.DB.Model(postSave).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()

	return postSave, err
}
