package database

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

type LikedPosts struct {
	DB *pg.DB
}

func (l *LikedPosts) CreatePostLike(postLike *model.LikedPost) (*model.LikedPost, error) {
	_, err := l.DB.Model(postLike).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		Returning("*").SelectOrInsert()
	return postLike, err
}

func (l *LikedPosts) GetLikedPostsByPostIDAndPagination(postID string, limit int, after string) (*model.LikedPosts, error) {
	var posts []*model.LikedPost
	var edges []*model.LikedPostsEdge
	var endCursor string

	query := l.DB.Model(&posts).Where("post_id = ?", postID).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")
	query.Limit(limit)

	if len(after) > 0 {
		query.Where("created_at < ?", after)
	}

	totalCount, err := query.SelectAndCount()

	for _, value := range posts {
		edges = append(edges, &model.LikedPostsEdge{Node: &model.LikedPost{
			ID:        value.ID,
			UserID:    value.UserID,
			PostID:    value.PostID,
			CreatedAt: value.CreatedAt,
		}})
	}

	if len(edges) > 0 {
		endCursor = util.ConvertCreateAtToCursor(*edges[len(edges)-1].Node.CreatedAt)
	}

	return &model.LikedPosts{
		TotalCount: &totalCount,
		Edges:      edges,
		PageInfo: &model.PageInfo{
			EndCursor: endCursor,
		},
	}, err

}

func (l *LikedPosts) GetPostLikeByUserIDAndPostID(userID, postID string) (*model.LikedPost, error) {
	var postLike model.LikedPost
	err := l.DB.Model(&postLike).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()

	return &postLike, err
}

func (l *LikedPosts) DeletePostLikeByID(id string) (*model.LikedPost, error) {
	DeletedAt := time.Now()
	var postLike = &model.LikedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := l.DB.Model(postLike).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return postLike, err
}
