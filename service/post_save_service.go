package service

import "context"

func (s *Service) CountPostTagsByTagID(ctx context.Context, id string) (int64, error) {
	return s.PostTags.CountPostTagsByTagID(ctx, id)
}
