package service

import (
	"context"

	"github.com/google/uuid"
)

func (s *Service) CountPostTagsByTagID(ctx context.Context, id uuid.UUID) (int64, error) {
	return s.PostTags.CountPostTagsByTagID(ctx, id)
}
