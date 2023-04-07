package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/plogto/core/db"
)

func (s *Service) GetPostAttachmentsByPostID(ctx context.Context, postID uuid.UUID) ([]*db.File, error) {
	postAttachments, _ := s.Files.GetFilesByPostID(ctx, postID)

	return postAttachments, nil
}
