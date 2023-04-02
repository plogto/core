package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) CreatePostMentions(ctx context.Context, userIDs []uuid.UUID, postID uuid.UUID) {
	for _, userID := range userIDs {
		s.PostMentions.CreatePostMention(ctx, userID, postID)
	}
}

func (s *Service) DeletePostMentions(ctx context.Context, userIDs []uuid.UUID, postID uuid.UUID) {
	for _, userID := range userIDs {
		s.PostMentions.DeletePostMention(ctx, userID, postID)
	}
}
