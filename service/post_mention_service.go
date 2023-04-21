package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) CreatePostMentions(ctx context.Context, userIDs []pgtype.UUID, postID pgtype.UUID) {
	for _, userID := range userIDs {
		s.PostMentions.CreatePostMention(ctx, userID, postID)
	}
}

func (s *Service) DeletePostMentions(ctx context.Context, userIDs []pgtype.UUID, postID pgtype.UUID) {
	for _, userID := range userIDs {
		s.PostMentions.DeletePostMention(ctx, userID, postID)
	}
}
