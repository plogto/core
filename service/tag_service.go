package service

import (
	"context"
	"errors"

	"github.com/favecode/poster-core/graph/model"
)

func (s *Service) SearchTag(ctx context.Context, expression string) (*model.Tags, error) {
	limit := 10
	page := 1
	tags, err := s.Tag.GetTagsByTagNameAndPagination(expression+"%", limit, page)

	if err != nil {
		return nil, errors.New("user not found")
	}

	return tags, nil
}
