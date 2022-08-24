package service

import "context"

func (s *Service) CountPostTagsByTagID(ctx context.Context, id string) (*int, error) {
	return s.PostTags.CountPostTagsByTagID(id)
}
