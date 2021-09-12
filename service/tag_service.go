package service

import (
	"context"
	"errors"

	"github.com/favecode/poster-core/config"
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

func (s *Service) CountTagByTagId(ctx context.Context, tagId string) (*int, error) {
	tagCount, _ := s.PostTag.CountPostTagsByTagId(tagId)

	return tagCount, nil
}

func (s *Service) GetTrends(ctx context.Context, input *model.PaginationInput) (*model.Tags, error) {
	var limit int = config.POSTS_PAGE_LIMIT
	var page int = 1

	if input != nil {
		if input.Limit != nil {
			limit = *input.Limit
		}

		if input.Page != nil && *input.Page > 0 {
			page = *input.Page
		}
	}
	tags, _ := s.PostTag.GetTagsOrderByCountTags(limit, page)

	return tags, nil
}
