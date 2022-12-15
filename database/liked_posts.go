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

func (l *LikedPosts) CreateLikedPost(likedPost *model.LikedPost) (*model.LikedPost, error) {
	_, err := l.DB.Model(likedPost).
		Where("user_id = ?user_id").
		Where("post_id = ?post_id").
		Where("deleted_at is ?", nil).
		SelectOrInsert()

	return likedPost, err
}

func (l *LikedPosts) GetLikedPostsByPostIDAndPageInfo(postID string, limit int, after string) (*model.LikedPosts, error) {
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

func (l *LikedPosts) GetLikedPostByUserIDAndPostID(userID, postID string) (*model.LikedPost, error) {
	var likedPost model.LikedPost
	err := l.DB.Model(&likedPost).Where("user_id = ?", userID).Where("post_id = ?", postID).Where("deleted_at is ?", nil).First()

	return &likedPost, err
}

func (l *LikedPosts) GetLikedPostByID(id string) (*model.LikedPost, error) {
	var likedPost model.LikedPost
	err := l.DB.Model(&likedPost).Where("id = ?", id).Where("deleted_at is ?", nil).First()

	return &likedPost, err
}

func (l *LikedPosts) DeleteLikedPostByID(id string) (*model.LikedPost, error) {
	DeletedAt := time.Now()
	var likedPost = &model.LikedPost{
		ID:        id,
		DeletedAt: &DeletedAt,
	}
	_, err := l.DB.Model(likedPost).Set("deleted_at = ?deleted_at").WherePK().Where("deleted_at is ?", nil).Returning("*").Update()
	return likedPost, err
}
