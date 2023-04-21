package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) CountPostTagsByTagID(ctx context.Context, id pgtype.UUID) (int64, error) {
	return s.PostTags.CountPostTagsByTagID(ctx, id)
}
