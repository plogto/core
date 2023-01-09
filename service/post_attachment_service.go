package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

func (s *Service) GetPostAttachmentsByPostID(ctx context.Context, postID string) ([]*db.File, error) {
	PostID, _ := uuid.Parse(postID)
	postAttachments, _ := s.Files.GetFilesByTicketMessageID(ctx, PostID)

	return postAttachments, nil
}
