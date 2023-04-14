package convertor

import (
	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/model"
)

func DBPostToModel(post *db.Post) *model.Post {
	return &model.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		ParentID:  post.ParentID,
		ChildID:   post.ChildID,
		Status:    model.PostStatus(post.Status),
		Content:   post.Content,
		Url:       post.Url,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func ModelPostToDB(post *model.Post) *db.Post {
	return &db.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		ParentID:  post.ParentID,
		ChildID:   post.ChildID,
		Status:    db.PostStatus(post.Status),
		Content:   post.Content,
		Url:       post.Url,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
