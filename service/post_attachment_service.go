package service

import (
	"context"

	"github.com/plogto/core/graph/model"
)

func (s *Service) GetPostAttachmentsByPostID(ctx context.Context, postID string) ([]*model.File, error) {
	postAttachments, _ := s.Files.GetFilesByPostId(postID)

	return postAttachments, nil
}
