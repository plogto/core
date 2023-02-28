package service

import (
	"context"
)

func (s *Service) CreatePostMentions(ctx context.Context, userIDs []string, postID string) {
	for _, userID := range userIDs {
		s.PostMentions.CreatePostMention(ctx, userID, postID)
	}
}

func (s *Service) DeletePostMentions(ctx context.Context, userIDs []string, postID string) {
	for _, userID := range userIDs {
		s.PostMentions.DeletePostMention(ctx, userID, postID)
	}
}
