package service

import (
	"time"

	"github.com/plogto/core/graph/model"
)

func (s *Service) CreatePostMentions(userIDs []string, postID string) {
	for _, userID := range userIDs {
		postMention := &model.PostMention{
			UserID: userID,
			PostID: postID,
		}
		s.PostMentions.CreatePostMention(postMention)
	}
}

func (s *Service) DeletePostMentions(userIDs []string, postID string) {
	for _, userID := range userIDs {
		postMention := &model.PostMention{
			UserID: userID,
			PostID: postID,
		}
		s.PostMentions.DeletePostMention(postMention)
	}
}

func (s *Service) DeletePostMentionsByPostID(postID string) {
	DeletedAt := time.Now()
	postMention := &model.PostMention{
		DeletedAt: &DeletedAt,
		PostID:    postID,
	}
	s.PostMentions.DeletePostMentionsByPostID(postMention)
}
