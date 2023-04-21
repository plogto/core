package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/plogto/core/db"
)

func (s *Service) GetPostAttachmentsByPostID(ctx context.Context, postID pgtype.UUID) ([]*db.File, error) {
	postAttachments, _ := s.Files.GetFilesByPostID(ctx, postID)

	return postAttachments, nil
}
